package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/pkg/node"
	"github.com/z7zmey/php-parser/pkg/node/stmt"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestInlineHtml(t *testing.T) {
	src := `<? ?> <div></div>`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    17,
		},
		Stmts: []node.Node{
			&stmt.Nop{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    5,
				},
			},
			&stmt.InlineHtml{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  5,
					EndPos:    17,
				},
				Value: " <div></div>",
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
