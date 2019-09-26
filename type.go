package exception

import (
	"fmt"
	"sync/atomic"
)

var nextId = int64(0)

type TriBool int

const (
	Undefined TriBool = iota
	True
	False
)

type Traits struct {
	HttpStatusCode int
	DisplayToUser  TriBool
	Description    string
	AllowSubclass  TriBool
	AllowThrow     TriBool
	Global         error
}

// Type defines an exception type
type Type struct {
	id             int64
	name           string
	parent         *Type
	httpStatusCode int
	display        bool
	description    string
	allowSubclass  bool
	allowThrow     bool
	global         error
}

func (t *Type) String() string {
	return fmt.Sprintf("%v: %v", t.name, t.description)
}

func (t *Type) Extend(name string, traits Traits) *Type {
	if !t.allowSubclass {
		panic(fmt.Errorf("extending a sealed type (%v)", t.description))
	}
	return newType(name, t, traits)
}

func (t *Type) extend(name string, traits Traits) *Type {
	return newType(name, t, traits)
}

func newType(name string, parent *Type, traits Traits) *Type {
	t := new(Type)
	t.name = name
	t.id = atomic.AddInt64(&nextId, 1)
	t.global = traits.Global
	t.parent = parent

	if traits.HttpStatusCode != 0 {
		t.httpStatusCode = traits.HttpStatusCode
	} else {
		t.httpStatusCode = parent.httpStatusCode
	}

	if traits.DisplayToUser == True {
		t.display = true
	} else if traits.DisplayToUser == False {
		t.display = false
	} else {
		t.display = parent.display
	}

	if traits.AllowSubclass == True {
		t.allowSubclass = true
	} else if traits.AllowSubclass == False {
		t.allowSubclass = false
	} else {
		t.allowSubclass = parent.allowSubclass
	}

	if traits.Description != "" {
		t.description = traits.Description
	} else {
		t.description = parent.description
	}

	if traits.AllowThrow == True {
		t.allowThrow = true
	} else if traits.AllowThrow == False {
		t.allowThrow = false
	} else {
		t.allowThrow = parent.allowThrow
	}

	return t
}

func (t *Type) IsSupertypeOfException(other Exception) bool {
	if other == nil {
		return false
	}

	o := other.Type()
	return t.IsSupertypeOf(o)
}

func (e *Type) IsSupertypeOf(other *Type) bool {
	for other != nil {
		if other.id == e.id {
			return true
		} else if other.id < e.id {
			return false
		}
		other = other.parent
	}

	return false
}
