//go:build js && wasm

package promise_test

import (
	"testing"

	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/promise"
	"github.com/stretchr/testify/assert"
)

func TestPromise(t *testing.T) {
	init, result := 1, 1
	ch := make(chan bool)

	promise.Resolve(init).Then(func(x js.Value) any {
		assert.Equal(t, init, x.Int())
		result += x.Int()
		return result
	}).Catch(func(x js.Value) any {
		assert.Equal(t, -init, x.Int())
		result += x.Int()
		return result
	}).Finally(func() any {
		t.Log("finally - channel")
		ch <- true
		return nil
	})

	<-ch

	assert.Equal(t, 2*init, result)

	result = init

	promise.Reject(-init).Then(func(x js.Value) any {
		assert.Equal(t, init, x.Int())
		result += x.Int()
		return result
	}).Catch(func(x js.Value) any {
		assert.Equal(t, -init, x.Int())
		result += x.Int()
		return result
	}).Finally(func() any {
		ch <- true
		return nil
	})

	<-ch

	assert.Equal(t, 0, result)
}

func TestPromiseAwait(t *testing.T) {
	assert.Equal(t, 1, promise.Resolve(1).Await().Int())
	assert.Equal(t, -1, promise.Reject(-1).Await().Int())

	assert.Equal(t, 2, promise.Resolve(1).Then(func(v js.Value) any {
		return v.Int() + 1
	}).Await().Int())

	assert.Equal(t, -2, promise.Reject(-1).Catch(func(v js.Value) any {
		return js.ValueOf(v.Int() - 1)
	}).Await().Int())

	result := promise.Resolve(map[string]any{"a": 1, "b": js.ValueOf(2)}).Then(func(v js.Value) any {
		v.Set("c", v.Get("a").Int()+v.Get("b").Int())
		return v
	}).Await()

	assert.Equal(t, 3, result.Get("c").Int())
}
