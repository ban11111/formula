package exp

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"reflect"
)

type IdentifierExpression struct {
	BaseExpression
	Name string
}

func (expression *IdentifierExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	return opt.NewArgumentWithType(expression.Name, reflect.String), nil
}

func NewIdentifierExpression(name string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &IdentifierExpression{
		Name: name,
	}

	return &result
}

type VarIdentifierExpression struct {
	BaseExpression
	Name string
}

//NewVarIdentifierExpression create new custom parameter which output like '[Parameter]'
func NewVarIdentifierExpression(name string) *opt.LogicalExpression {
	var result opt.LogicalExpression = &IdentifierExpression{
		Name: name[1 : len(name)-2],
	}

	return &result
}

func (expression *VarIdentifierExpression) Evaluate(context *opt.FormulaContext) (*opt.Argument, error) {
	if p, ok := expression.Context.Parameters[expression.Name]; ok {
		return opt.NewArgument(p), nil
	}

	return nil, fmt.Errorf("variable %s can not be resolved", expression.Name)
}
