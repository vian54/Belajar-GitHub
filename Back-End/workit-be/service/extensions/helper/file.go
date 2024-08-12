package helper

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func GenerateExcel(ctx context.Context, data interface{}, filePath, sheetName string) (err error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err = errors.New(fmt.Sprintln("no data to be generated", r))
	// 	}
	// }()

	xlsx := excelize.NewFile()
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	sliceValue := reflect.ValueOf(data)
	if sliceValue.Kind() != reflect.Slice {
		err = errors.New("Data Should be slice")
		return
	}

	dataValue := sliceValue.Index(0)
	if dataValue.Kind() != reflect.Struct {
		err = errors.New("Data should be slice of struct")
		return
	}

	structType := dataValue.Type()

	for i := 0; i < sliceValue.Len(); i++ {
		structValue := sliceValue.Index(i)
		for j := 0; j < structValue.NumField(); j++ {
			field := structType.Field(j)
			xlsxField := field.Tag.Get("xlsxField")
			if xlsxField == "" {
				continue
			}

			xlsxFieldSplit := strings.Split(xlsxField, ":")
			if len(xlsxFieldSplit) != 2 {
				err = errors.New("xlsxField tag should be in <columnAlphabetic>:<columnName> format")
				return
			}

			xlsxColumnChar := strings.ToUpper(strings.TrimSpace(xlsxFieldSplit[0]))

			if i == 0 {
				xlsxColumnName := strings.TrimSpace(xlsxFieldSplit[1])
				xlsx.SetCellValue(sheetName, fmt.Sprintf("%s%d", xlsxColumnChar, i+1), xlsxColumnName)
			}

			fieldValue := structValue.Field(j).Interface()
			val := fmt.Sprintf("%v", fieldValue)
			if IsPointer(fieldValue) {
				dt := GetValueFromPointer(fieldValue)
				if dt == nil {
					continue
				}
				val = fmt.Sprintf("%v", dt)
			}
			xlsx.SetCellValue(sheetName, fmt.Sprintf("%s%d", xlsxColumnChar, i+2), val)
		}
	}

	err = xlsx.SaveAs(filePath)
	if err != nil {
		return
	}

	return
}

func GetDocBase64(ctx context.Context, filePath string) (res string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	defer os.Remove(filePath)
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	res = base64.StdEncoding.EncodeToString(fileContent)
	return
}
