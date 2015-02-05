package filter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func newRansackParam(param interface{}, kind string) *ransakParam {
	ellipsisRx := regexp.MustCompile(`^[\d]+..[\d]+$`)

	arrayRx := regexp.MustCompile(`^\[[\d|,]*\]$`)

	rparam := &ransakParam{
		value:      param,
		kind:       kind,
		ellipsisRx: ellipsisRx,
		arrayRx:    arrayRx,
	}

	rparam.findStrRepresentation()

	return rparam
}

type ransakParam struct {
	value             interface{}
	kind              string
	StrRepresentation string
	ellipsisRx        *regexp.Regexp
	arrayRx           *regexp.Regexp
}

func (this *ransakParam) findStrRepresentation() {
	paramStr := fmt.Sprintf("%v", this.value)

	if str, isEllipsis := this.getFromEllipsis(paramStr); isEllipsis {
		this.StrRepresentation = str
		return
	}

	if str, isArray := this.getFromArray(paramStr); isArray {
		this.StrRepresentation = str
		return
	}

	this.StrRepresentation = paramStr
}

func (this *ransakParam) getFromEllipsis(param string) (string, bool) {
	if this.ellipsisRx.MatchString(param) {
		values := strings.Split(param, "..")

		//if it wasn't a number, the regexp would have failed,
		//so we can omit the error
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])

		var strValues []string

		for i := start; i <= end; i++ {
			strValues = append(strValues, strconv.Itoa(i))
		}

		return strings.Join(strValues, ","), true
	}
	return "", false
}

func (this *ransakParam) getFromArray(param string) (string, bool) {
	if this.arrayRx.MatchString(param) {
		r := regexp.MustCompile(`[\[|\]]`)

		param = r.ReplaceAllString(param, "")

		return strings.Trim(param, ","), true
	}
	return "", false
}
