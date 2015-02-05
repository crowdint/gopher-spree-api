package repositories

import (
	"fmt"
	"reflect"
	"strings"
)

type OperatorFunction func(re *RansakEmulator)

type Node struct {
	Name  string
	Nodes []*Node
	Apply OperatorFunction
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
}

func (this *RansakEmulator) ToSql(input string, param interface{}) string {
	this.reset()

	this.tokenize(input)

	this.paramKind = reflect.TypeOf(param).String()

	for this.pos = 0; this.pos < len(this.toEvaluate); this.pos++ {
		token := this.toEvaluate[this.pos]

		if node, isCandidate := isCandidateToOperator(token); isCandidate {
			if foundNode, found := this.find(node, this.pos); found {

				foundNode.Apply(this)

			} else {

				this.evaluated(token)

			}
		} else {

			this.evaluated(token)

		}
	}

	this.replaceValue(param)

	return strings.Trim(this.template, " ")
}

func (this *RansakEmulator) reset() {
	this.toEvaluate = []string{}
	this.evaluatedTokens = []string{}
	this.template = ""
	this.paramKind = ""
}

func (this *RansakEmulator) tokenize(input string) {
	this.toEvaluate = strings.Split(input, this.separator)
}

func (this *RansakEmulator) find(nodeParam *Node, pos int) (*Node, bool) {
	if pos >= len(this.toEvaluate) {
		return nil, false
	}

	next := this.toEvaluate[pos]

	if nodeParam.Name != next {
		return nil, false
	}

	if len(nodeParam.Nodes) > 0 {
		for _, node := range nodeParam.Nodes {
			if foundNode, found := this.find(node, pos+1); found {
				return foundNode, true
			}
		}
	} else {
		this.pos = pos
		return nodeParam, true
	}

	return nil, false
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

func isCandidateToOperator(item string) (*Node, bool) {
	for _, node := range Tree.Nodes {
		if node.Name == item {
			return node, true
		}
	}
	return &Node{}, false
}

var Tree = &Node{
	Name: "Operators",
	Nodes: []*Node{
		&Node{
			Name: "or",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.template += "OR "
			},
		},
		&Node{
			Name: "and",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.template += "AND "
			},
		},
		&Node{
			Name: "eq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "matches",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.placeholder + "'")
			},
		},
		&Node{
			Name: "cont",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.placeholder + "%'")
			},
		},
		&Node{
			Name: "lt",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("< " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "lteq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("<= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "gt",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("> " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "gteq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder(">= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "start",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.placeholder + "%'")
			},
		},
		&Node{
			Name: "end",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.placeholder + "'")
			},
		},
		&Node{
			Name: "not",
			Nodes: []*Node{
				&Node{
					Name: "eq",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("<> " + re.getCorrectSqlFormat(re.placeholder))
					},
				},
				&Node{
					Name: "in",
					Apply: func(re *RansakEmulator) {
						re.appendField()
					},
				},
				&Node{
					Name: "cont",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '%" + re.placeholder + "%'")
					},
				},
				&Node{
					Name: "start",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '" + re.placeholder + "%'")
					},
				},
			},
		},
		&Node{
			Name: "does",
			Nodes: []*Node{
				&Node{
					Name: "not",
					Nodes: []*Node{
						&Node{
							Name: "match",
							Apply: func(re *RansakEmulator) {
								re.appendField()
								re.replacePlaceholder("NOT LIKE '" + re.placeholder + "'")
							},
						},
					},
				},
			},
		},
	},
}
