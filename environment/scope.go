package environment

import (
  "gossip/types"
)

type Scope struct {
  Functions map[string] func([]types.LispElement) types.LispElement
  Variables map[string] types.LispElement
  Parent *Scope
}

func (scope *Scope) LookupFunction(label string) (func([]types.LispElement) types.LispElement, bool) {
  ret, ret_exists := scope.Functions[label]
  if ret_exists {
    return ret, true
  }
  if scope.Parent != nil {
    return scope.Parent.LookupFunction(label)
  }
  return nil, false
}
func (scope *Scope) LookupVariable(label string) (types.LispElement, bool) {
  ret, ret_exists := scope.Variables[label]
  if ret_exists {
    return ret, true
  }
  if scope.Parent != nil {
    return scope.Parent.LookupVariable(label)
  }
  return nil, false
}
func (scope *Scope) AddVariable(label string, value types.LispElement) {
  scope.Variables[label] = value
}

