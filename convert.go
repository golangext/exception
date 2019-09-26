package exception

func convert(v interface{}) Exception {
	if v == nil {
		return nil
	}

	if e, ok := v.(Exception); ok {
		return e
	}

	var t *Type

	t = knownGlobalErrObjects.Get(v)

	if t == nil {
		t = knownGlobalErrTypes.Get(v)
	}

	if t == nil {
		t = Panic
	}

	if err, ok := v.(error); ok {
		return t.new(err, "%v", err.Error())
	}

	return t.new(nil, "%v", v)
}
