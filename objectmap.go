package exception

import (
	"runtime"
	"unsafe"
)

var knownGlobalErrObjects = make(objectMap)
var knownGlobalErrTypes = make(typeMap)

func (typeMap) id(v interface{}) uintptr {
	type iface struct {
		typePtr uintptr
		dataPtr uintptr
	}

	ifPtr := (*iface)(unsafe.Pointer(&v))
	key := ifPtr.typePtr
	runtime.KeepAlive(v)
	return key
}

func (objectMap) id(v interface{}) uintptr {
	type iface struct {
		typePtr uintptr
		dataPtr uintptr
	}

	ifPtr := (*iface)(unsafe.Pointer(&v))
	key := ifPtr.dataPtr
	runtime.KeepAlive(v)
	return key
}

type objectMap map[uintptr]*Type

func (a objectMap) Add(err error, t *Type) {
	key := a.id(err)
	a[key] = t
}

func (a objectMap) Get(err interface{}) *Type {
	key := a.id(err)
	if v, ok := a[key]; ok {
		return v
	}
	return nil
}

type typeMap map[uintptr]*Type

func (a typeMap) Add(err error, t *Type) {
	key := a.id(err)
	a[key] = t
}

func (a typeMap) Get(err interface{}) *Type {
	key := a.id(err)
	if v, ok := a[key]; ok {
		return v
	}
	return nil
}

func (parent *Type) extendErrObj(name string, err error) *Type {
	t := parent.extend(name, Traits{
		AllowSubclass: False,
		AllowThrow:    True,
		Global:        err,
	})

	knownGlobalErrObjects.Add(err, t)
	return t
}

func (parent *Type) extendErrType(name string, err error) *Type {
	t := parent.extend("", Traits{
		AllowSubclass: False,
		AllowThrow:    True,
	})

	knownGlobalErrTypes.Add(err, t)
	return t
}
