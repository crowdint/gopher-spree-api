package utils

import (
	"reflect"
)

func Collect(collection interface{}, fieldName string) (result []interface{}) {
	slice := reflect.ValueOf(collection)
	if slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if field := slice.Index(i).FieldByName(fieldName); field.IsValid() {
				value := field.Interface()
				result = append(result, value)
			}
		}
	}
	return
}

func ToMap(collection interface{}, key string, multiple bool) map[int64]interface{} {
	slice := reflect.ValueOf(collection)
	result := make(map[int64]interface{})

	if slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if field := slice.Index(i).FieldByName(key); field.IsValid() {
				key := field.Int()

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
		}
	}

	return result
}
