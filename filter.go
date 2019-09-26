package exception

type Filter struct {
	m map[int64]*Type
}

func FilterOf(t ...*Type) *Filter {
	f := &Filter{make(map[int64]*Type)}
	for _, v := range t {
		cur := v
		for cur != nil {
			f.m[cur.id] = cur
			cur = cur.parent
		}
	}
	return f
}

func (f *Filter) Catch(fn func(e Exception)) {
	if v := recover(); v != nil {
		f.catch(v, fn)
	}
}

func (f *Filter) catch(v interface{}, fn func(e Exception)) {
	e := convert(v)

	t := e.Type()

	for t != nil {
		if _, ok := f.m[t.id]; ok {
			fn(e)
			return
		}
		t = t.parent
	}

	// No match, rethrow the original value
	rethrow(v)
}
