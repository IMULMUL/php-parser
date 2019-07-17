package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/pkg/node/name"
	"github.com/z7zmey/php-parser/pkg/position"

	"github.com/z7zmey/php-parser/pkg/node/expr"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/pkg/node"
	"github.com/z7zmey/php-parser/pkg/node/stmt"
)

func TestClassConstFetch(t *testing.T) {
	src := `<? Foo::Bar;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    12,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    12,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    11,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    6,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    6,
								},
								Value: "Foo",
							},
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    11,
						},
						Value: "Bar",
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src))
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestStaticClassConstFetch(t *testing.T) {
	src := `<? static::bar;`

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
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    14,
					},
					Class: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    9,
						},
						Value: "static",
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    14,
						},
						Value: "bar",
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src))
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src))
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
