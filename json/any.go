//go:build js && wasm

package json

import (
	"fmt"
	"time"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/errors"
)

// -----------------------------------------------------------------------------

func returnTypeError(v js.Value, typ string) error {
	return errors.New(fmt.Sprintf("%s is not a %s", v.Type(), typ), builtin.TypeError)
}

// -----------------------------------------------------------------------------

func TimeToString(tm time.Time) string {
	return tm.Format(TimeFormat)
}

// -----------------------------------------------------------------------------

func ToIntE(v js.Value) (int, error) {
	if v.Type() == js.TypeNumber {
		return v.Int(), nil
	}
	return 0, returnTypeError(v, "number")
}

// -----------------------------------------------------------------------------

func ToIntDefault(v js.Value, def int) int {
	if ret, err := ToIntE(v); err == nil {
		return ret
	} else {
		return def
	}
}

// -----------------------------------------------------------------------------

func ToInt(v js.Value) (ret int) {
	ret, _ = ToIntE(v)
	return
}

// -----------------------------------------------------------------------------

func ToFloat64E(v js.Value) (float64, error) {
	if v.Type() == js.TypeNumber {
		return v.Float(), nil
	}
	return 0.0, returnTypeError(v, "number")
}

// -----------------------------------------------------------------------------

func ToFloat64Default(v js.Value, def float64) float64 {
	if ret, err := ToFloat64E(v); err == nil {
		return ret
	} else {
		return def
	}
}

// -----------------------------------------------------------------------------

func ToFloat(v js.Value) (ret float64) {
	ret, _ = ToFloat64E(v)
	return
}

// -----------------------------------------------------------------------------

func ToStringE(v js.Value) (string, error) {
	if v.Type() == js.TypeString {
		return v.String(), nil
	}
	return "", returnTypeError(v, "string")
}

// -----------------------------------------------------------------------------

func ToString(v js.Value) (ret string) {
	ret, _ = ToStringE(v)
	return
}

// -----------------------------------------------------------------------------

func ToStringDefault(v js.Value, def string) string {
	if ret, err := ToStringE(v); err == nil {
		return ret
	} else {
		return def
	}
}

// -----------------------------------------------------------------------------

func ToBoolE(v js.Value) (bool, error) {
	if v.Type() == js.TypeBoolean {
		return v.Bool(), nil
	}
	return false, returnTypeError(v, "boolean")
}

// -----------------------------------------------------------------------------

func ToBoolDefault(v js.Value, def bool) bool {
	if ret, err := ToBoolE(v); err == nil {
		return ret
	} else {
		return def
	}
}

// -----------------------------------------------------------------------------

func ToBool(v js.Value) (ret bool) {
	ret, _ = ToBoolE(v)
	return
}

// -----------------------------------------------------------------------------

func ToTimeE(v js.Value) (time.Time, error) {
	if v.Type() == js.TypeString {
		if tm, err := time.Parse(TimeFormat, v.String()); err != nil {
			return time.Time{}, errors.TypeError(err)
		} else {
			return tm, nil
		}
	}
	return time.Time{}, returnTypeError(v, "string")
}

// -----------------------------------------------------------------------------

func ToTimeDefault(v js.Value, def time.Time) time.Time {
	if ret, err := ToTimeE(v); err == nil {
		return ret
	} else {
		return def
	}
}

// -----------------------------------------------------------------------------

func ToTime(v js.Value) (ret time.Time) {
	ret, _ = ToTimeE(v)
	return
}

// -----------------------------------------------------------------------------

type Any js.Value

// -----------------------------------------------------------------------------

func (a Any) GetIntE(name string) (int, error) {
	return ToIntE(js.Value(a).Get(name))
}

func (a Any) GetIntDefault(name string, def int) int {
	return ToIntDefault(js.Value(a).Get(name), def)
}

// -----------------------------------------------------------------------------

func (a Any) GetInt(name string) (ret int) {
	ret, _ = a.GetIntE(name)
	return
}

// -----------------------------------------------------------------------------

func (a Any) GetFloatE(name string) (float64, error) {
	return ToFloat64E(js.Value(a).Get(name))
}

// -----------------------------------------------------------------------------

func (a Any) GetFloatDefault(name string, def float64) float64 {
	return ToFloat64Default(js.Value(a).Get(name), def)
}

// -----------------------------------------------------------------------------

func (a Any) GetFloat(name string) (ret float64) {
	ret, _ = a.GetFloatE(name)
	return
}

// -----------------------------------------------------------------------------

func (a Any) GetStringE(name string) (string, error) {
	return ToStringE(js.Value(a).Get(name))
}

// -----------------------------------------------------------------------------

func (a Any) GetStringDefault(name string, def string) string {
	return ToStringDefault(js.Value(a).Get(name), def)
}

// -----------------------------------------------------------------------------

func (a Any) GetString(name string) (ret string) {
	ret, _ = a.GetStringE(name)
	return
}

// -----------------------------------------------------------------------------

func (a Any) GetBoolE(name string) (bool, error) {
	return ToBoolE(js.Value(a).Get(name))
}

// -----------------------------------------------------------------------------

func (a Any) GetBoolDefault(name string, def bool) bool {
	return ToBoolDefault(js.Value(a).Get(name), def)
}

// -----------------------------------------------------------------------------

func (a Any) GetBool(name string) (ret bool) {
	ret, _ = a.GetBoolE(name)
	return
}

// -----------------------------------------------------------------------------

func (a Any) GetTimeE(name string) (time.Time, error) {
	return ToTimeE(js.Value(a).Get(name))
}

// -----------------------------------------------------------------------------

func (a Any) GetTimeDefault(name string, def time.Time) time.Time {
	return ToTimeDefault(js.Value(a).Get(name), def)
}

// -----------------------------------------------------------------------------

func (a Any) GetTime(name string) (ret time.Time) {
	ret, _ = a.GetTimeE(name)
	return
}

// -----------------------------------------------------------------------------

type Result struct {
	js.Value
	lastErr error
}

// -----------------------------------------------------------------------------

func (r *Result) Error() error {
	return r.lastErr
}

// -----------------------------------------------------------------------------

func (r *Result) IntVar(a *int, name string) *Result {
	if r.lastErr == nil {
		*a, r.lastErr = ToIntE(r.Value.Get(name))
	}
	return r
}

// -----------------------------------------------------------------------------

func (r *Result) Float64Var(a *float64, name string) *Result {
	if r.lastErr == nil {
		*a, r.lastErr = ToFloat64E(r.Value.Get(name))
	}
	return r
}

// -----------------------------------------------------------------------------

func (r *Result) StringVar(a *string, name string) *Result {
	if r.lastErr == nil {
		*a, r.lastErr = ToStringE(r.Value.Get(name))
	}
	return r
}

// -----------------------------------------------------------------------------

func (r *Result) BoolVar(a *bool, name string) *Result {
	if r.lastErr == nil {
		*a, r.lastErr = ToBoolE(r.Value.Get(name))
	}
	return r
}

// -----------------------------------------------------------------------------

func (r *Result) TimeVar(a *time.Time, name string) *Result {
	if r.lastErr == nil {
		*a, r.lastErr = ToTimeE(r.Value.Get(name))
	}
	return r
}

// -----------------------------------------------------------------------------

func ResultOf(v js.Value) *Result {
	return &Result{
		Value:   v,
		lastErr: nil,
	}
}
