package utils

import (
	"reflect"
)

func Collect(collection interface{}, field string) (result []interface{}) {
	slice := reflect.ValueOf(collection)

	for i := 0; i < slice.Len(); i++ {
		value := slice.Index(i).FieldByName(field).Interface()
		result = append(result, value)
	}
	return
}

func ToMap(collection interface{}, key string, multiple bool) map[int64]interface{} {
	slice := reflect.ValueOf(collection)

	result := make(map[int64]interface{})

	for i := 0; i < slice.Len(); i++ {
		key := slice.Index(i).FieldByName(key).Int()

		if multiple {
			newValue := slice.Index(i)

			if result[key] == nil {
				result[key] = []interface{}{newValue.Interface()}
			} else {
				result[key] = append(result[key].([]interface{}), newValue.Interface())
			}

		} else {
			result[key] = slice.Index(i).Interface()
		}
	}

	return result
}
