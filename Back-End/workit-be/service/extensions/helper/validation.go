package helper

import (
	"reflect"

	"github.com/thoas/go-funk"
)

func IsStruct(val interface{}) bool {
	// Use reflection to get the type of the value
	valueType := reflect.TypeOf(val)

	// Check if the type is a struct
	return valueType.Kind() == reflect.Struct
}

func IsStructOrPointerToStruct(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	kind := valueType.Kind()

	return kind == reflect.Struct || (kind == reflect.Ptr && valueType.Elem().Kind() == reflect.Struct)
}

func IsPointer(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Pointer
}

func IsPointerOfStruct(i interface{}) bool {
	if !IsPointer(i) {
		return false
	}

	return reflect.TypeOf(i).Elem().Kind() == reflect.Struct
}
func IsPointerOfInt(i interface{}) bool {
	if !IsPointer(i) {
		return false
	}

	return reflect.TypeOf(i).Elem().Kind() == reflect.Int
}

func IsMap(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Map
}

func IsSlice(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Slice
}

func SliceContains(collections interface{}, element interface{}) bool {
	return funk.Contains(collections, element)
}
