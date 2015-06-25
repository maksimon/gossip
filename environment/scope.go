package environment

import (
  "gossip/types"
)

type Scope struct {
  functions map[string] func([]types.LispElement) types.LispElement
  variables map[string] types.LispElement
  Parent *Scope
}

func (scope *Scope) LookupFunction(label string) (func([]types.LispElement) types.LispElement, bool) {
  ret, ret_exists := scope.functions[label]
  if ret_exists {
    return ret, true
  }
  if scope.Parent != nil {
    return scope.Parent.LookupFunction(label)
  }
  return nil, false
}
func (scope *Scope) LookupVariable(label string) (types.LispElement, bool) {
  ret, ret_exists := scope.variables[label]
  if ret_exists {
    return ret, true
  }
  if scope.Parent != nil {
    return scope.Parent.LookupVariable(label)
  }
  return nil, false
}
func (scope *Scope) AddVariable(label string, value types.LispElement) {
  scope.variables[label] = value
}

func (scope *Scope) AddFunction(label string, value func([]types.LispElement) types.LispElement) {
  scope.functions[label] = value 
}
