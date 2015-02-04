package repositories

import (
	"fmt"
	"reflect"
	"strings"
)

type OperatorFunction func(re *RansakEmulator)

type Node struct {
	Name     string
	Nodes    []Node
	Function OperatorFunction
}

func NewRansakEmulator() *RansakEmulator {
	return &RansakEmulator{
		separator:   "_",
		placeholder: "{{.}}",
	}
}

type RansakEmulator struct {
	toEvaluate      []string
	evaluatedTokens []string
	template        string
	separator       string
	placeholder     string
	paramKind       string
	pos             int
	currentOperator []string
}

func (this *RansakEmulator) ToSql(input string, param interface{}) string {
	this.reset()

	this.tokenize(input)

	this.paramKind = reflect.TypeOf(param).String()

	for this.pos = 0; this.pos < len(this.toEvaluate); this.pos++ {
		token := this.toEvaluate[this.pos]

		if node, matched := isOperator(token); matched {
			if !this.find(node, this.pos) {
				this.evaluated(token)
			}
		} else {
			this.evaluated(token)
		}
	}

	this.replaceValue(param)

	return this.template
}

func (this *RansakEmulator) reset() {
	this.toEvaluate = []string{}
	this.evaluatedTokens = []string{}
	this.currentOperator = []string{}
	this.template = ""
	this.paramKind = ""
}

func (this *RansakEmulator) tokenize(input string) {
	this.toEvaluate = strings.Split(input, this.separator)
}

func (this *RansakEmulator) find(nodeParam Node, pos int) bool {
	if pos >= len(this.toEvaluate) {
		return false
	}

	next := this.toEvaluate[pos]

	if nodeParam.Name != next {
		return false
	}

	if len(nodeParam.Nodes) > 0 {
		for _, node := range nodeParam.Nodes {
			if this.find(node, pos+1) {
				return true
			}
		}
	} else {
		this.pos = pos
		nodeParam.Function(this)
		return true
	}

	return false
}

func (this *RansakEmulator) appendField() {
	field := strings.Join(this.evaluatedTokens, this.separator)
	this.evaluatedTokens = []string{}
	this.template += field + " {{.}} "
}

func (this *RansakEmulator) evaluated(token string) {
	this.evaluatedTokens = append(this.evaluatedTokens, token)
}

func (this *RansakEmulator) replacePlaceholder(replaceFor string) {
	this.template = strings.Replace(
		this.template,
		this.placeholder,
		replaceFor,
		-1,
	)
}

func (this *RansakEmulator) replaceValue(value interface{}) {
	strValue := fmt.Sprintf("%v", value)

	this.replacePlaceholder(strValue)
}

func (this *RansakEmulator) getCorrectSqlFormat(value string) string {
	if this.paramKind == "string" {
		return "'" + value + "'"
	}
	return value
}

func isOperator(item string) (Node, bool) {
	for _, node := range Tree.Nodes {
		if node.Name == item {
			return node, true
		}
	}
	return Node{}, false
}

var Tree = Node{
	Name: "Operators",
	Nodes: []Node{
		Node{
			Name: "or",
			Function: func(re *RansakEmulator) {
				re.appendField()
				re.template += "OR "
			},
		},
		Node{
			Name: "and",
			Function: func(re *RansakEmulator) {
				re.appendField()
				re.template += "AND "
			},
		},
		Node{
			Name: "eq",
			Function: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		Node{
			Name: "matches",
			Function: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.placeholder + "'")
			},
		},
		Node{
			Name: "cont",
			Function: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.placeholder + "%'")
			},
		},
		Node{
			Name: "not",
			Nodes: []Node{
				{
					Name: "eq",
					Function: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("<> " + re.getCorrectSqlFormat(re.placeholder))
					},
				},
				{
					Name: "in",
					Function: func(re *RansakEmulator) {
						re.appendField()
					},
				},
			},
		},
	},
}
