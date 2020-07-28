// +build !debug

package purescript

func Apply(f Any, args ...Any) Any {
	result := f
	for _, arg := range args {
		fn, _ := result.(Fn)
		result = fn(arg)
	}
	return result
}

func Run(f Any, args ...Any) Any {
	fn, _ := f.(EffFn)
	return fn()
}

func Contains(dict Any, key string) bool {
	d, _ := dict.(Dict)
	_, found := d[key]
	return found
}

func Length(arr Any) Any {
	a, _ := arr.([]Any)
	return len(a)
}

func Index(arr Any, idx Any) Any {
	a, _ := arr.([]Any)
	i, _ := idx.(int)
	return a[i]
}

func CopyDict(dict Any) Any {
	d, _ := dict.(Dict)
	cpy := make(Dict)
	for k, v := range d {
		cpy[k] = v
	}
	return cpy
}
