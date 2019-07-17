package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/pkg/node"
	"github.com/z7zmey/php-parser/pkg/node/expr"
	"github.com/z7zmey/php-parser/pkg/node/expr/assign"
	"github.com/z7zmey/php-parser/pkg/node/stmt"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestShortList(t *testing.T) {
	src := `<? [$a] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    13,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    13,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    12,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    7,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    6,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    6,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  4,
											EndPos:    6,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    12,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    12,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestShortListArrayIndex(t *testing.T) {
	src := `<? [$a[]] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    15,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    15,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    9,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    8,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    8,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  4,
											EndPos:    6,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  4,
												EndPos:    6,
											},
											Value: "a",
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    14,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    14,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestShortListList(t *testing.T) {
	src := `<? [list($a)] = $b;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    19,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    19,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    18,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    13,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  4,
									EndPos:    12,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  4,
										EndPos:    12,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  9,
												EndPos:    11,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  9,
													EndPos:    11,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 1,
														EndLine:   1,
														StartPos:  9,
														EndPos:    11,
													},
													Value: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  16,
							EndPos:    18,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    18,
							},
							Value: "b",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
