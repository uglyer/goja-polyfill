package fetch

import (
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"log"
	"net/http"
	"net/url"
)

type Fetch struct {
}

type Request struct {
	Body     string            `json:"body"`
	Headers  map[string]string `json:"headers"`
	Method   string            `json:"method"`
	Redirect string            `json:"redirect"`
}

type Response struct {
	Status     int32  `json:"status"`
	StatusText string `json:"statusText"`
	OK         bool   `json:"ok"`
	Redirected bool   `json:"redirected"`
	URL        string `json:"url"`
}

//
func (f *Fetch) prepareReq(rawURL string, data map[string]any) (*http.Request, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("url '%s' is not valid", rawURL))
	}
	//req, err := http.NewRequest(jsReq.Method, url.String(), body)
	//var jsReq Request
	//if len(args) > 1 {
	//	if !args[1].IsObject() {
	//		return nil, errors.New("2nd argument must be an object")
	//	}
	//	reader := strings.NewReader(args[1].JSONStringify())
	//	if err := json.NewDecoder(reader).Decode(&jsReq); err != nil {
	//		return nil, err
	//	}
	//}
	//
	//if jsReq.Method != "" {
	//	jsReq.Method = "GET"
	//}
	//
	//var body io.Reader
	//if jsReq.Method != "GET" {
	//	body = strings.NewReader(jsReq.Body)
	//}
	//
	//req, err := http.NewRequest(jsReq.Method, url.String(), body)
	//if err != nil {
	//	return nil, err
	//}
	//for k, v := range jsReq.Headers {
	//	headerName := http.CanonicalHeaderKey(k)
	//	req.Header.Set(headerName, v)
	//}
	//
	//if req.Header.Get("Accept") == "" {
	//	req.Header.Set("Accept", "*/*")
	//}
	//
	//if req.Header.Get("Connection") == "" {
	//	req.Header.Set("Connection", "close")
	//}
	//
	//req.Header.Set("Redirect", jsReq.Redirect)
	//
	//return req, nil
}

func (f *Fetch) fetch(url string, data map[string]any) {
	log.Println(url, data)
}

func (f *Fetch) Inject(vm *goja.Runtime) error {
	if err := vm.Set("fetch", f.fetch); err != nil {
		return err
	}
	return nil
}

func Inject(vm *goja.Runtime) error {
	fetch := &Fetch{}
	return fetch.Inject(vm)
}
