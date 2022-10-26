package polyfill

import "github.com/dop251/goja"

type Polyfill interface {
	Inject(vm *goja.Runtime) error
}
