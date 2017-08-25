package is

import (
	"testing"
)

// Is ...
type Is interface {
	Nil(interface{})
	OK(interface{})
}

type is struct {
	t        *testing.T
	isStrict bool
	text     string
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

func (i *is) fail() {
	i.t.Logf(i.text)
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
		i.fail()
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
		i.fail()
	}
}
