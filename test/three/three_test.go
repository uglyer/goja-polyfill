package three

import (
	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/uglyer/goja-polyfill/polyfill/console"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func newVM(t *testing.T) *goja.Runtime {
	vm := goja.New()
	assert.NoError(t, console.Inject(vm))
	return vm
}

func loadJS(t *testing.T, url string, filename string) []byte {
	_, err := os.Stat(filename)
	if err == nil {
		b, err := ioutil.ReadFile(filename)
		assert.NoError(t, err)
		return b
	}
	resp, err := http.Get(url)
	assert.NoError(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	err = ioutil.WriteFile(filename, body, 0600)
	assert.NoError(t, err)
	return body
}

func TestThee(t *testing.T) {
	vm := newVM(t)
	js := string(loadJS(t, "https://three.com.com.sb/build/three.js", "./dist/three.js"))
	_, err := vm.RunString(js)
	assert.NoError(t, err)
	scene, err := vm.RunString("new THREE.Scene()")
	assert.NoError(t, err)
	log.Println(scene)
}
