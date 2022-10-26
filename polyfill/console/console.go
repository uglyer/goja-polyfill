package console

import (
	"github.com/dop251/goja"
	"log"
)

type Console struct {
}

func (c *Console) log(args ...any) {
	log.Println(args...)
}

func (c *Console) Polyfill(vm *goja.Runtime) error {
	consoleObj := vm.NewObject()
	if err := consoleObj.Set("log", c.log); err != nil {
		return err
	}
	if err := vm.Set("console", consoleObj); err != nil {
		return err
	}
	return nil
}

func Inject(vm *goja.Runtime) error {
	console := &Console{}
	return console.Polyfill(vm)
}
