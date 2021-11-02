package complier

import "testing"

func TestEvaluate(t *testing.T) {
	// !(10 + 2 * 4)
	ast := &Ast{
		kind: AstKindNot,
		children: []*Ast{
			{
				kind: AstKindPlus,
				children: []*Ast{
					{
						kind:  AstKindValue,
						value: 10,
					},
					{
						kind: AstKindMul,
						children: []*Ast{
							{
								kind:  AstKindValue,
								value: 2,
							},
							{
								kind:  AstKindValue,
								value: 4,
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
