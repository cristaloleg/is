package is

import (
	"math"
	"reflect"
	"regexp"
	"strings"
)

// Is is a main interface in this lib
type Is interface {
	Nil(interface{})
	OK(interface{})
	Err(error)
	NoErr(error)

	Type(interface{}, interface{})
	Impl(impl interface{}, object interface{})
	Match(string, string)
	Pos(interface{})
	Neg(interface{})
	Zero(interface{})
	Int(interface{})
	Float(interface{})
	NaN(interface{})
	Empty(interface{})
	Closed(interface{})
	Contains(interface{}, interface{})
	Equal(interface{}, interface{})
	NotEqual(interface{}, interface{})
}

// T interface for testing.T
type T interface {
	Log(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Helper()
}

type is struct {
	t        T
	isStrict bool
	text     string
}

// New ...
func New(t T) Is {
	is := &is{
		t:        t,
		isStrict: false,
	}
	return is
}

// NewStrict ...
func NewStrict(t T) Is {
	is := &is{
		t:        t,
		isStrict: true,
	}
	return is
}

func (i *is) fail() {
	i.t.Log(i.text)
	i.t.FailNow()
}

// Nil ...
func (i *is) Nil(obj interface{}) {
	i.t.Helper()
	if obj == nil {
		return
	}
	i.t.Errorf("expected nil, but got %v", obj)
	if i.isStrict {
		i.fail()
	}
}

// OK ...
func (i *is) OK(obj interface{}) {
	i.t.Helper()
	if obj == nil {
		i.t.Errorf("unexpected  nil")
		if i.isStrict {
			i.fail()
		}
		return
	}
	switch value := obj.(type) {
	case string:
		if value == "" {
			i.t.Errorf("unexpected empty string")
		}
	case bool:
		if !value {
			i.t.Errorf("unexpected false")
		}
	}

	if obj == 0 {
		i.t.Errorf("unexpected zero")
	}
}

// Err ...
func (i *is) Err(err error) {
	if err != nil {
		return
	}
	i.t.Errorf("expected error, but got nil")
	if i.isStrict {
		i.fail()
	}
}

// NoErr ...
func (i *is) NoErr(err error) {
	if err == nil {
		return
	}
	i.t.Errorf("unexpected error %v", err)
	if i.isStrict {
		i.fail()
	}
}

// Type ...
func (i *is) Type(ttype interface{}, object interface{}) {
	x := reflect.TypeOf(ttype)
	y := reflect.TypeOf(object)
	if x == y {
		return
	}
	i.t.Errorf("expected type %v, but got %v", x, y)
	if i.isStrict {
		i.fail()
	}
}

func (i *is) Impl(impl interface{}, object interface{}) {}
func (i *is) Match(pattern string, text string) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		i.t.FailNow()
		return
	}
	if !r.MatchString(text) {
		i.t.FailNow()
	}
}

func (i *is) Pos(obj interface{}) {
	switch value := obj.(type) {
	case byte:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
		if value < 0 {
			if i.isStrict {
				i.fail()
			}
		}
	}
}
func (i *is) Neg(obj interface{}) {
	switch value := obj.(type) {
	case byte:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
		if value > 0 {
			if i.isStrict {
				i.fail()
			}
		}
	}
}

func (i *is) Zero(obj interface{}) {
	switch value := obj.(type) {
	case byte:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
		if value != 0 {
			if i.isStrict {
				i.fail()
			}
		}
	}
}
func (i *is) Int(obj interface{}) {
	switch obj.(type) {
	default:
		if i.isStrict {
			i.fail()
		}
	case byte:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint16:
	case uint32:
	case uint64:
		//
	}
}

func (i *is) Float(obj interface{}) {
	switch obj.(type) {
	default:
		if i.isStrict {
			i.fail()
		}
	case float32:
	case float64:
		//
	}
}

func (i *is) NaN(obj interface{}) {
	switch value := obj.(type) {
	case float32:
	case float64:
		if !math.IsNaN(value) {
			if i.isStrict {
				i.fail()
			}
		}
	default:
	}
}
func (i *is) Empty(obj interface{}) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if value.Len() != 0 {
			if i.isStrict {
				i.fail()
			}
		}
	default:
	}
}

func (i *is) Closed(obj interface{}) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Chan:
		_, ok := value.Recv()
		if ok {
			if i.isStrict {
				i.fail()
			}
		}
	default:
	}
}

func (i *is) Contains(container interface{}, element interface{}) {
	value := reflect.ValueOf(container)
	switch value.Kind() {
	case reflect.Array:

	case reflect.Map:
	case reflect.Slice:
	case reflect.String:
		s := container.(string)
		e, ok := element.(string)
		if !ok || !strings.Contains(s, e) {
			if i.isStrict {
				i.fail()
			}
		}
	default:
	}
}

func (i *is) Equal(interface{}, interface{})    {}
func (i *is) NotEqual(interface{}, interface{}) {}
