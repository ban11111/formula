package formula

import (
	"github.com/ban11111/formula/internal/cache"
	"github.com/ban11111/formula/opt"
)

//Register custom function which implement opt.Function
func Register(fs ...opt.Function) error {
	for _, f := range fs {
		if err := cache.Register(f); err != nil {
			return err
		}
	}
	return nil
}

//RegisterGlobalParameter register global parameter which will be used in all the runtime
//func RegisterGlobalParameter(name string, value interface{}) {
//
//}
