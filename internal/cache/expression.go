package cache

import "github.com/ban11111/formula/opt"

var defaultExpression = make(map[string]*opt.LogicalExpression, 512)

func Store(option opt.Option, originalExpression string, logicalExpression opt.LogicalExpression) {

}

func Restore(option opt.Option, originalExpression string) *opt.LogicalExpression {
	return nil
}
