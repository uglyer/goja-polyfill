package console

import (
	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newVM(t *testing.T) *goja.Runtime {
	vm := goja.New()
	err := Inject(vm)
	assert.NoError(t, err)
	return vm
}

func TestLog(t *testing.T) {
	vm := newVM(t)
	_, err := vm.RunString(`console.log(1,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.trace(2,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.debug(3,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.info(4,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.warn(5,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.error(6,"test2",{"test":"ttt"},["1","2",3])`)
	assert.NoError(t, err)
}

func TestTimer(t *testing.T) {
	vm := newVM(t)
	_, err := vm.RunString(`console.time("timer1")`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.time("timer1")`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.time("timer2")`)
	assert.NoError(t, err)
	time.Sleep(time.Duration(1) * time.Millisecond)
	_, err = vm.RunString(`console.timeEnd("timer1")`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.timeEnd("timer1")`)
	assert.NoError(t, err)
	_, err = vm.RunString(`console.timeEnd("timer2")`)
	assert.NoError(t, err)
}
