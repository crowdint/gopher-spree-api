package repositories

import (
	//"log"
	"strings"
)

var operators = []string{
	"cont",
	"or",
}

var (
	current       []string
	template      string
	separationStr = "_"
)

func lex(input string) {
	items := strings.Split(input, separationStr)

	for _, item := range items {
		if isOperator(item) {
			if len(current) > 0 {
				joinedStr := strings.Join(current, separationStr)
				appendField(joinedStr)
				current = []string{}
			}
			appendOperator(item)
		} else {
			current = append(current, item)
		}
	}
}

func isOperator(item string) bool {
	for _, op := range operators {
		if op == item {
			return true
		}
	}
	return false
}

func appendField(item string) {
	template += (item + " {{.}} ")
}

func appendOperator(operator string) {
	switch operator {
	case "or":
		template += "OR "
	case "cont":
		template = strings.Replace(template, "{{.}}", "LIKE %{{.}}%", -1)
	}
}
