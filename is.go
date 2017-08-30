package is

import "reflect"

// Is is a main interface in this lib
type Is interface {
	Nil(interface{})
	OK(interface{})
	Err(error)
	NoErr(error)

	Type(interface{}, interface{})
	Impl(impl interface{}, object interface{})
	Match(string, interface{})
	Pos(interface{})
	Neg(interface{})
	Zero(interface{})
	Int(interface{})
	Float(interface{})
	NaN(interface{})
	Empty(interface{})
	Closed(interface{})
	Filled(interface{})
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

// Nil ...
func (i *is) OK(obj interface{}) {
	i.t.Helper()
	if obj != nil {
		return
	}
	i.t.Errorf("expected not nil")
	if i.isStrict {
		i.fail()
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
	i.t.Errorf("expected nil, but got %v", err)
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
func (i *is) Match(string, interface{})                 {}
func (i *is) Pos(interface{})                           {}
func (i *is) Neg(interface{})                           {}
func (i *is) Zero(interface{})                          {}
func (i *is) Int(interface{})                           {}
func (i *is) Float(interface{})                         {}
func (i *is) NaN(interface{})                           {}
func (i *is) Empty(interface{})                         {}
func (i *is) Closed(interface{})                        {}
func (i *is) Filled(interface{})                        {}
func (i *is) Contains(interface{}, interface{})         {}
func (i *is) Equal(interface{}, interface{})            {}
func (i *is) NotEqual(interface{}, interface{})         {}
