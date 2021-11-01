package complier

import "testing"

func TestEvaluate(t *testing.T) {
	// !(10 + 2 * 4)
	ast := &Ast{
		Kind: AstKindNot,
		Children: []*Ast{
			{
				Kind: AstKindPlus,
				Children: []*Ast{
					{
						Kind:  AstKindValue,
						Value: 10,
					},
					{
						Kind: AstKindMul,
						Children: []*Ast{
							{
								Kind:  AstKindValue,
								Value: 2,
							},
							{
								Kind:  AstKindValue,
								Value: 4,
							},
						},
					},
				},
			},
		},
	}

	ret, err := ast.Evaluate()
	t.Logf("%#v\n", err)
	t.Logf("%#v\n", ret)

}
