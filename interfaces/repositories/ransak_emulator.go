package repositories

import (
	"fmt"
	"reflect"
	"strings"
)

var ransakOperators = []string{
	"cont",
	"or",
	"and",
	"eq",
	"matches",
}

type RansakEmulator struct {
	current                []string
	template               string
	separator              string
	placeholder            string
	evaluatingMultiTokenOp bool
}

func NewRansakEmulator() *RansakEmulator {
	return &RansakEmulator{
		current:     []string{},
		separator:   "_",
		placeholder: "{{.}}",
	}
}

func (this *RansakEmulator) ToSql(input string, param interface{}) string {
	this.reset()

	items := strings.Split(input, this.separator)

	kind := reflect.TypeOf(param).String()

	for _, item := range items {
		this.appendToTemplate(item, kind)
	}

	this.replaceValue(param)

	return this.template
}

func (this *RansakEmulator) reset() {
	this.current = []string{}
	this.template = ""
}

func (this *RansakEmulator) appendToTemplate(item string, kind string) {
	if isOperator(item) {
		this.appendCurrentField()
		this.appendOperator(item, kind)
	} else {
		this.current = append(this.current, item)
	}
}

func (this *RansakEmulator) replaceValue(param interface{}) {
	paramStr := fmt.Sprintf("%v", param)

	this.replacePlaceholder(paramStr)
}

func (this *RansakEmulator) appendCurrentField() {
	if len(this.current) > 0 {
		joinedStr := strings.Join(this.current, this.separator)
		this.template += (joinedStr + " " + this.placeholder + " ")
		this.current = []string{}
	}
}

func (this *RansakEmulator) appendOperator(operator string, kind string) {
	switch operator {
	case "or":
		this.template += "OR "
	case "and":
		this.template += "AND "
	case "cont":
		this.replacePlaceholder("LIKE '%" + this.placeholder + "%'")
	case "eq":
		replaceFor := "= " + this.getAsSqlType(this.placeholder, kind)

		this.replacePlaceholder(replaceFor)
	case "matches":
		this.replacePlaceholder("LIKE " + this.getAsSqlType(this.placeholder, kind))
	}
}

func (this *RansakEmulator) getAsSqlType(param string, kind string) string {
	if kind == "string" {
		return "'" + param + "'"
	}

	return param
}

func (this *RansakEmulator) replacePlaceholder(replaceFor string) {
	this.template = strings.Replace(
		this.template,
		this.placeholder,
		replaceFor, -1)
}

func isOperator(item string) bool {
	for _, op := range ransakOperators {
		if op == item {
			return true
		}
	}
	return false
}
