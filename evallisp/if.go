package evallisp

import (
  "gossip/environment"
  "gossip/types"
)

func If(if_args []types.LispElement, scope *environment.Scope) types.LispElement {
  if len(if_args) < 2 {
    panic("Not enough arguments to 'if'")
  }
  if !types.IsNil(eval(*scope, if_args[0])) {
    return eval(*scope, if_args[1])
  }
  if len(if_args) > 2 {
    return eval(*scope, if_args[2])
  }
  return types.NewList()
}
