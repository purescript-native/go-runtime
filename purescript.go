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

const Undefined = "\x00<undefined>\x00"

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
