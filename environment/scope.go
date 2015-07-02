package environment

import (
  "gossip/types"
)

type Scope struct {
  Variables map[string] types.LispElement
  Parent *Scope
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

func(scope *Scope) LookupFunction(label string) (*types.LispFunction, bool) {
  ret, ret_exists := scope.LookupVariable(label)
  return ret.(*types.LispFunction), ret_exists
}

func (scope *Scope) AddVariable(label string, value types.LispElement) {
  scope.Variables[label] = value
}
