package json

import (
	"strconv"
)

func ToS(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}
