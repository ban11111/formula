package fs

import "github.com/yidane/formula/opt"

type IIFFunction struct {
}

func (*IIFFunction) Name() string {
	return "iif"
}

func (*IIFFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewIIFFunction() *IIFFunction {
	return &IIFFunction{}
}