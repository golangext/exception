package exception

import "fmt"

type Exception interface {
	error
	Type() *Type
	Cause() error
	PrintStackTrace()
	Is(*Type) bool
}

type exception struct {
	message string
	t       *Type
	cause   error
}

func (e *exception) Is(t *Type) bool {
	return t.IsSupertypeOfException(e)
}

func (e *exception) PrintStackTrace() {
}

func (e *exception) Error() string {
	return e.message
}

func (e *exception) Type() *Type {
	return e.t
}

func (e *exception) Cause() error {
	return e.cause
}

func (t *Type) new(cause error, messageFormat string, args ...interface{}) Exception {
	e := &exception{message: fmt.Sprintf(messageFormat, args...), t: t, cause: cause}
	return e
}

func (t *Type) New(cause error, messageFormat string, args ...interface{}) Exception {
	if !t.allowThrow {
		e := &exception{message: "New called on virtual type", t: uncatchable, cause: nil}
		panic(e)
	}
	return t.new(cause, messageFormat, args...)
}

func Throw(e Exception) {
	if !e.Type().allowThrow {
		e = &exception{message: "Throw called on virtual type", t: uncatchable, cause: nil}
	}
	panic(e)
}

func ThrowOnException(e Exception) {
	if e != nil {
		Throw(e)
	}
}

func (t *Type) ThrowNew(cause error, messageFormat string, args ...interface{}) {
	e := t.New(cause, messageFormat, args...)
	Throw(e)
}

func (t *Type) ThrowNewOnErr(cause error, messageFormat string, args ...interface{}) {
	if cause != nil {
		t.ThrowNew(cause, messageFormat, args...)
	}
}

func (t *Type) Catch(fn func(e Exception)) {
	if v := recover(); v != nil {
		t.catch(v, fn)
	}
}

func (t *Type) catch(v interface{}, fn func(e Exception)) {
	e := convert(v)

	if t.IsSupertypeOfException(e) {
		fn(e)
		return
	}

	// No match, rethrow the original value
	panic(v)
}
