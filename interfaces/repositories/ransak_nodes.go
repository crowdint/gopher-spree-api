package repositories

type OperatorFunction func(re *RansakEmulator)

type Node struct {
	Name  string
	Nodes []*Node
	Apply OperatorFunction
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
