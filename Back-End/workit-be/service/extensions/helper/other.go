package helper

import "reflect"

func GetValueFromPointer(pointer interface{}) interface{} {
	if pointer == nil {
		return nil
	}

	value := reflect.ValueOf(pointer)

	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = value.Elem()

		if value.CanInterface() {
			return value.Interface()
		}
	}
	return nil
}
