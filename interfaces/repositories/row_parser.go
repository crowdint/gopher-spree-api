package repositories

import (
	"database/sql"
	"reflect"
)

func ParseRow(row []interface{}, target interface{}) {
	s := reflect.ValueOf(target).Elem()
	for i := 0; i < s.NumField(); i++ {
		if row[i] == nil {
			continue
		}

		f := s.Field(i)

		strVal := reflect.ValueOf(row[i]).Convert(f.Type())

		f.Set(strVal)
	}
}

func ParseAllRows(kind interface{}, rows *sql.Rows) ([]interface{}, error) {
	arr := make([]interface{}, 0)

	cols, err := rows.Columns()
	if err != nil {
		return arr, err
	}

	rawResult := make([]interface{}, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice

	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	s := reflect.ValueOf(kind).Elem()

	for rows.Next() {
		rows.Scan(dest...)

		newElement := reflect.New(s.Type()).Elem()

		for i := 0; i < newElement.NumField(); i++ {
			if rawResult[i] == nil {
				continue
			}

			f := newElement.Field(i)

			strVal := reflect.ValueOf(rawResult[i]).Convert(f.Type())

			f.Set(strVal)
		}

		arr = append(arr, newElement.Interface())
	}

	return arr, nil
}
