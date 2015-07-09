package environment

import (
  "gossip/types"
  "fmt"
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

func (scope *Scope) SetVariable(label string, value types.LispElement) {
  _, var_exists := scope.Variables[label]
  if var_exists {
    scope.Variables[label] = value
  } else if (scope.Parent != nil) {
    scope.Parent.SetVariable(label, value)
  } else {
    panic(fmt.Sprintf("variable with label %s does not exist", label))
  }
}

func (scope *Scope) AddVariable(label string, value types.LispElement) {
  _, var_exists := scope.LookupVariable(label)
  if !var_exists {
    scope.Variables[label] = value
  }
}
