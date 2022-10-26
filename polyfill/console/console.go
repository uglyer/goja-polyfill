package console

import (
	"fmt"
	"github.com/dop251/goja"
)

type Console struct {
}

func (c *Console) log(args ...any) {
	fmt.Println(args...)
}

func (c *Console) Polyfill(vm *goja.Runtime) error {
	consoleObj := vm.NewObject()
	if err := consoleObj.Set("log", c.log); err != nil {
		return err
	}
	if err := consoleObj.Set("trace", c.log); err != nil {
		return err
	}
	if err := consoleObj.Set("debug", c.log); err != nil {
		return err
	}
	if err := consoleObj.Set("info", c.log); err != nil {
		return err
	}
	if err := consoleObj.Set("warn", c.log); err != nil {
		return err
	}
	if err := consoleObj.Set("error", c.log); err != nil {
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
