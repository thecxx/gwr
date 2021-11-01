package complier

import (
	"errors"

	"github.com/spf13/cast"
)

type AstKind uint8

const (
	AstKindUnknown AstKind = iota
	AstKindValue
	AstKindPlus
	AstKindMinus
	AstKindMul
	AstKindDiv
	AstKindAnd
	AstKindOr
	AstKindNot
)

type Ast struct {
	Kind     AstKind
	Value    interface{}
	Children []*Ast
}

func (a *Ast) Evaluate() (ret interface{}, err error) {
	var (
		op1 interface{}
		op2 interface{}
	)

	switch a.Kind {
	case AstKindPlus:
		if op1, err = a.Children[0].Evaluate(); err != nil {
			return nil, err
		}
		if op2, err = a.Children[1].Evaluate(); err != nil {
			return nil, err
		}
		ret = OperatePlus(op1, op2)
	case AstKindMinus:
		if op1, err = a.Children[0].Evaluate(); err != nil {
			return nil, err
		}
		if op2, err = a.Children[1].Evaluate(); err != nil {
			return nil, err
		}
		ret = OperateMinus(op1, op2)
	case AstKindMul:
		if op1, err = a.Children[0].Evaluate(); err != nil {
			return nil, err
		}
		if op2, err = a.Children[1].Evaluate(); err != nil {
			return nil, err
		}
		ret = OperateMul(op1, op2)
	case AstKindDiv:
		if op1, err = a.Children[0].Evaluate(); err != nil {
			return nil, err
		}
		if op2, err = a.Children[1].Evaluate(); err != nil {
			return nil, err
		}
		ret = OperateDiv(op1, op2)
	case AstKindNot:
		if op1, err = a.Children[0].Evaluate(); err != nil {
			return nil, err
		}
		ret = !cast.ToBool(op1)
	case AstKindValue:
		ret = OperateValue(a.Value)
	case AstKindUnknown:
		fallthrough
	default:
		err = errors.New("unsupported kind")
	}

	return
}

func OperateValue(op1 interface{}) interface{} {
	return op1
}

func OperatePlus(op1, op2 interface{}) int {
	return cast.ToInt(op1) + cast.ToInt(op2)
}

func OperateMinus(op1, op2 interface{}) int {
	return cast.ToInt(op1) + cast.ToInt(op2)
}

func OperateMul(op1, op2 interface{}) int {
	return cast.ToInt(op1) * cast.ToInt(op2)
}

func OperateDiv(op1, op2 interface{}) int {
	return cast.ToInt(op1) / cast.ToInt(op2)
}
