package console

import (
	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newVM(t *testing.T) *goja.Runtime {
	vm := goja.New()
	err := Inject(vm)
	assert.NoError(t, err)
	return vm
}

func TestLog(t *testing.T) {
	vm := newVM(t)
	_, err := vm.RunString(`console.log("test","test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
}
