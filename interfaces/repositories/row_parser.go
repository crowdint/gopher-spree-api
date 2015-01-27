package repositories

import (
	"reflect"
)

func ParseRow(row []interface{}, target interface{}) {
	s := reflect.ValueOf(target).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		strVal := reflect.ValueOf(row[i])

		f.Set(strVal)
	}
}
