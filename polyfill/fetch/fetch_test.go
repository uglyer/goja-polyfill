package fetch

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

func TestFetchGet(t *testing.T) {
	vm := newVM(t)
	_, err := vm.RunString(`fetch("https://www.baidu.com",{header:{"test":"test"}})`)
	assert.NoError(t, err)
}
