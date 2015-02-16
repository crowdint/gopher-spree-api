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
