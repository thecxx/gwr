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
	kind     AstKind
	value    interface{}
	children []*Ast
}

func (a *Ast) Evaluate() (ret interface{}, err error) {
	var (
		op1 interface{}
		op2 interface{}
	)

	switch a.kind {
	case AstKindPlus:
		op1, op2, err = a.EvaluateOp1_2()
		if err != nil {
			return nil, err
		}
		ret = OperatePlus(op1, op2)
	case AstKindMinus:
		op1, op2, err = a.EvaluateOp1_2()
		if err != nil {
			return nil, err
		}
		ret = OperateMinus(op1, op2)
	case AstKindMul:
		op1, op2, err = a.EvaluateOp1_2()
		if err != nil {
			return nil, err
		}
		ret = OperateMul(op1, op2)
	case AstKindDiv:
		op1, op2, err = a.EvaluateOp1_2()
		if err != nil {
			return nil, err
		}
		ret = OperateDiv(op1, op2)
	case AstKindNot:
		op1, err = a.EvaluateOp1()
		if err != nil {
			return nil, err
		}
		ret = !cast.ToBool(op1)
	case AstKindValue:
		ret = OperateValue(a.value)
	case AstKindUnknown:
		fallthrough
	default:
		err = errors.New("unsupported kind")
	}

	return
}

func (a *Ast) EvaluateOp1() (op1 interface{}, err error) {
	if op1, err = a.children[0].Evaluate(); err != nil {
		return nil, err
	}
	return
}

func (a *Ast) EvaluateOp1_2() (op1, op2 interface{}, err error) {
	if op1, err = a.children[0].Evaluate(); err != nil {
		return nil, nil, err
	}
	if op2, err = a.children[1].Evaluate(); err != nil {
		return nil, nil, err
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
