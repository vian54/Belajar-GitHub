package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(data interface{}) (res []byte) {
	res, _ = json.Marshal(data)
	return
}

func MapToJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func JsonStringToMap(data string) (res map[string]interface{}) {
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		return
	}
	return
}

func StructToMap(input interface{}, tag ...string) map[string]interface{} {
	usedTag := "json"
	if len(tag) > 0 {
		usedTag = tag[0]
	}

	result := make(map[string]interface{})

	// Use reflection to get the value of the struct
	valueType := reflect.ValueOf(input)
	if valueType.Kind() == reflect.Ptr {
		// If it's a pointer, dereference it to get the struct value
		valueType = valueType.Elem()
	}

	if valueType.Kind() == reflect.Struct {
		valueTypeType := valueType.Type()

		for i := 0; i < valueType.NumField(); i++ {
			field := valueType.Field(i)
			fieldName := valueTypeType.Field(i).Name

			// Try to get the JSON tag; if it exists, use it as the key
			tag := valueTypeType.Field(i).Tag.Get(usedTag)
			if tag != "" {
				result[tag] = field.Interface()
			} else {
				// Fall back to using the field name as the key
				result[fieldName] = field.Interface()
			}
		}
	}

	return result
}

func StructToMapString(input interface{}, tag ...string) map[string]string {
	usedTag := "json"
	if len(tag) > 0 {
		usedTag = tag[0]
	}

	result := make(map[string]string)

	// Use reflection to get the value of the struct
	valueType := reflect.ValueOf(input)
	if valueType.Kind() == reflect.Ptr {
		// If it's a pointer, dereference it to get the struct value
		valueType = valueType.Elem()
	}

	if valueType.Kind() == reflect.Struct {
		valueTypeType := valueType.Type()

		for i := 0; i < valueType.NumField(); i++ {
			field := valueType.Field(i)
			fieldName := valueTypeType.Field(i).Name

			// Try to get the JSON tag; if it exists, use it as the key
			tag := valueTypeType.Field(i).Tag.Get(usedTag)
			if tag != "" {
				result[tag] = fmt.Sprintf("%s", field.Interface())
			} else {
				// Fall back to using the field name as the key
				result[fieldName] = fmt.Sprintf("%s", field.Interface())
			}
		}
	}

	return result
}

func MapAnyToStruct(data map[string]interface{}, resultStruct interface{}, structTag ...string) error {
	resultValue := reflect.ValueOf(resultStruct)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Struct {
		return errors.New("resultStruct must be a pointer to a struct")
	}

	tag := "json"
	if len(structTag) > 0 {
		tag = structTag[0]
	}

	structType := resultValue.Elem().Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tagValue := field.Tag.Get(tag)
		if tagValue == "" {
			// If the field doesn't have a JSON tag, skip it
			continue
		}
		fieldName := field.Name
		if value, ok := data[tagValue]; ok {
			structFieldValue := resultValue.Elem().Field(i)
			mapValue := reflect.ValueOf(value)
			if mapValue.Type().AssignableTo(field.Type) {
				structFieldValue.Set(mapValue)
			} else {
				return errors.New(fmt.Sprintf("field %s has an incompatible type in the map", fieldName))
			}
		}
	}
	return nil
}

func StripLeadingZerosAndDecimal(input string) string {
	for i, char := range input {
		if char != '0' {
			return input[i:]
		}
	}
	return ""
}

func StringFloatToFloat(inputStr string) (res float64) {
	strippedStr := StripLeadingZerosAndDecimal(inputStr)

	// Parse the remaining string as an integer
	res, err := strconv.ParseFloat(strippedStr, 64)
	if err != nil {
		return
	}

	return
}

func StringToInt(inputStr string) (res int64) {
	strippedStr := StripLeadingZerosAndDecimal(inputStr)

	res, err := strconv.ParseInt(strippedStr, 10, 64)
	if err != nil {
		return
	}

	return
}

func BoolToInt(payload interface{}) int8 {
	switch value := payload.(type) {
	case string:
		if strings.ToUpper(value) == "TRUE" {
			return 1
		}
	case bool:
		if value {
			return 1
		}
	}
	return 0
}
