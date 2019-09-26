package exception

import "unsafe"

var knownGlobalErrObjects = make(objectMap)
var knownGlobalErrTypes = make(typeMap)

func (typeMap) id(v interface{}) uintptr {
	type iface struct {
		typePtr uintptr
		dataPtr uintptr
	}

	ifPtr := (*iface)(unsafe.Pointer(&v))
	key := ifPtr.typePtr
	return key
}

func (objectMap) id(v interface{}) uintptr {
	type iface struct {
		typePtr uintptr
		dataPtr uintptr
	}

	ifPtr := (*iface)(unsafe.Pointer(&v))
	key := ifPtr.dataPtr
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

func knownGlobalObject(err error, parent *Type) *Type {
	t := parent.extend("", Traits{
		AllowSubclass: False,
		AllowThrow:    True,
		Global:        err,
	})

	knownGlobalErrObjects.Add(err, t)
	return t
}

func knownGlobalType(err error, parent *Type) *Type {
	t := parent.extend("", Traits{
		AllowSubclass: False,
		AllowThrow:    True,
		Global:        err,
	})

	knownGlobalErrTypes.Add(err, t)
	return t
}
