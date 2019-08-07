package ajax

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var typeProtoMsg = reflect.TypeOf((*proto.Message)(nil)).Elem()
var typeError = reflect.TypeOf((*error)(nil)).Elem()

type response struct {
	code    int
	result  []byte
	lastErr error
}

func (resp *response) String() string {
	result := ""
	if resp.result != nil {
		result = base64.StdEncoding.EncodeToString(resp.result)
	}

	err := ""
	if resp.lastErr != nil {
		err = resp.lastErr.Error()
	}

	return fmt.Sprintf(`{"code":%d, "result":"%s", "error":"%s"}`, resp.code, result, err)
}

type callback struct {
	types []reflect.Type
	fn    reflect.Value
}

func (cb *callback) invoke(resp *response) {
	args := make([]reflect.Value, 0, len(cb.types))

	for _, t := range cb.types {

		if t.Implements(typeProtoMsg) {
			v := reflect.New(t.Elem())
			resp.lastErr = proto.Unmarshal(resp.result, v.Interface().(proto.Message))
			if resp.lastErr != nil {
				return
			}
			args = append(args, v)
			continue
		}

		if t.Implements(typeError) {
			args = append(args, reflect.ValueOf(resp.lastErr))
			continue
		}

		switch t.Kind() {
		case reflect.Int:
			args = append(args, reflect.ValueOf(resp.code))
		case reflect.String:
			args = append(args, reflect.ValueOf(string(resp.result)))
		case reflect.Slice:
			args = append(args, reflect.ValueOf(resp.result))
		case reflect.Ptr:
			if t.Elem().Kind() == reflect.Struct {
				v := reflect.New(t.Elem())
				resp.lastErr = json.Unmarshal(resp.result, v.Interface())
				if resp.lastErr != nil {
					return
				}
				args = append(args, v)
			}
		}
	}

	cb.fn.Call(args)
}

func checkCB(fn interface{}) *callback {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return nil
	}

	size := t.NumIn()
	if size > 3 {
		return nil
	}

	types := make([]reflect.Type, size, size)
	for i := 0; i < size; i++ {
		types[i] = t.In(i)
	}

	return &callback{
		types: types,
		fn:    reflect.ValueOf(fn),
	}
}

// Client ...
type Client struct {
	resp     *response
	finished bool
	done     *callback
	fail     *callback
	always   *callback
}

func (cli *Client) callDone() {
	if !cli.finished || cli.done == nil {
		return
	}
	if cli.resp.code >= http.StatusOK && cli.resp.code < http.StatusMultipleChoices {
		cli.done.invoke(cli.resp)
	}
}

func (cli *Client) callFail() {
	if !cli.finished || cli.fail == nil {
		return
	}

	if cli.resp.code >= http.StatusBadRequest {
		cli.fail.invoke(cli.resp)
	}
}

func (cli *Client) callAlways() {
	if !cli.finished || cli.always == nil {
		return
	}

	cli.always.invoke(cli.resp)
}

// Done ...
func (cli *Client) Done(fn interface{}) *Client {
	if cb := checkCB(fn); cb != nil {
		cli.done = cb
		cli.callDone()
	}
	return cli
}

// Fail ...
func (cli *Client) Fail(fn interface{}) *Client {
	if cb := checkCB(fn); cb != nil {
		cli.fail = cb
		cli.callFail()
	}
	return cli
}

// Always ...
func (cli *Client) Always(fn interface{}) *Client {
	if cb := checkCB(fn); cb != nil {
		cli.always = cb
		cli.callAlways()
	}
	return cli
}

func (cli *Client) Error() string {
	if cli.resp.lastErr == nil {
		return ""
	}
	return cli.resp.lastErr.Error()
}

func (cli *Client) String() string {
	if cli.resp != nil {
		return cli.resp.String()
	}

	return ""
}

// Call ...
func call(method, url string, data ...interface{}) *Client {
	cli := &Client{resp: &response{}}

	var req *http.Request

	if len(data) > 0 {
		dataBytes, err := json.Marshal(data[0])
		if err != nil {
			cli.resp.lastErr = err
			return cli
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(dataBytes))
		if err != nil {
			cli.resp.lastErr = err
			return cli
		}
		req.Header.Add("Content-Type", "application/json")
	} else {
		var err error
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			cli.resp.lastErr = err
			return cli
		}
	}

	go func(c *Client, r *http.Request) {
		defer func() {
			c.finished = true
			c.callDone()
			c.callFail()
			c.callAlways()
		}()

		xcli := new(http.Client)
		resp, err := xcli.Do(r)

		if err != nil {
			c.resp.lastErr = err
			return
		}

		c.resp.code = resp.StatusCode
		c.resp.result, c.resp.lastErr = ioutil.ReadAll(resp.Body)
	}(cli, req)

	return cli
}

// Get ...
func Get(url string, data ...interface{}) *Client {
	return call("GET", url, data...)
}

// Post ...
func Post(url string, data ...interface{}) *Client {
	return call("POST", url, data...)
}

// Put ...
func Put(url string, data ...interface{}) *Client {
	return call("PUT", url, data...)
}

// Delete ...
func Delete(url string, data ...interface{}) *Client {
	return call("DELETE", url, data...)
}
