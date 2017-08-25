package is

import (
	"testing"
)

// Is ...
type Is interface {
	Fail(string)
	Nil(interface{})
	OK(interface{})
}

type is struct {
	t        *testing.T
	isStrict bool
}

// New ...
func New(t *testing.T) Is {
	is := &is{
		t:        t,
		isStrict: false,
	}
	return is
}

// NewStrict ...
func NewStrict(t *testing.T) Is {
	is := &is{
		t:        t,
		isStrict: true,
	}
	return is
}

// Fail ...
func (i *is) Fail(text string) {
	i.t.Logf(text)
	i.t.FailNow()
}

// Nil ...
func (i *is) Nil(object interface{}) {
	i.t.Helper()
	if object == nil {
		return
	}
	i.t.Errorf("expected nil %v", object)
	if i.isStrict {
		i.Fail("")
	}
}

// Nil ...
func (i *is) OK(object interface{}) {
	i.t.Helper()
	if object != nil {
		return
	}
	i.t.Errorf("expected not nil %v", object)
	if i.isStrict {
		i.Fail("")
	}
}
