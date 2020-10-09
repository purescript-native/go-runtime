package purescript

import "sync"

type Any = interface{}

type Fn = func(Any) Any
type EffFn = func() Any
type Dict = map[string]Any

type Once = sync.Once

type Fn1 = func(Any) Any
type Fn2 = func(Any, Any) Any
type Fn3 = func(Any, Any, Any) Any
type Fn4 = func(Any, Any, Any, Any) Any
type Fn5 = func(Any, Any, Any, Any, Any) Any
type Fn6 = func(Any, Any, Any, Any, Any, Any) Any
type Fn7 = func(Any, Any, Any, Any, Any, Any, Any) Any
type Fn8 = func(Any, Any, Any, Any, Any, Any, Any, Any) Any
type Fn9 = func(Any, Any, Any, Any, Any, Any, Any, Any, Any) Any
type Fn10 = func(Any, Any, Any, Any, Any, Any, Any, Any, Any, Any) Any

var undefined = "<undefined>"
var Undefined Any = &undefined

var foreign = make(map[string]Dict)

func Foreign(moduleName string) Dict {
	moduleDict, found := foreign[moduleName]
	if !found {
		moduleDict = Dict{"module": moduleName}
		foreign[moduleName] = moduleDict
	}
	return moduleDict
}

func Get(dict Dict, key string) Any {
	value, ok := dict[key]
	if !ok {
		panic("Foreign value '" + key + "' for purescript module '" + dict["module"].(string) + "' not found")
	}
	return value
}

func Apply(f Any, args ...Any) Any {
	result := f
	for _, arg := range args {
		fn := result.(Fn)
		result = fn(arg)
	}
	return result
}

func Run(f Any, args ...Any) Any {
	fn := f.(EffFn)
	return fn()
}

func Contains(dict Any, key string) bool {
	d := dict.(Dict)
	_, found := d[key]
	return found
}

func Is(dict Any, key string) bool {
	d := dict.(Dict)
	_, found := d[key]
	return found
}

func Length(arr Any) Any {
	a := arr.([]Any)
	return len(a)
}

func Index(arr Any, idx Any) Any {
	a := arr.([]Any)
	i := idx.(int)
	return a[i]
}

func CopyDict(dict Any) Any {
	d := dict.(Dict)
	cpy := make(Dict)
	for k, v := range d {
		cpy[k] = v
	}
	return cpy
}
