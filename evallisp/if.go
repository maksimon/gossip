package evallisp

import (
  "gossip/types"
)

func If(if_args []types.LispElement) types.LispElement {
  if len(if_args) < 2 {
    panic("Not enough arguments to 'if'")
  }
  if !types.IsNil(if_args[0]) {
    return if_args[1]
  }
  if len(if_args) >= 2 {
    return if_args[2]
  }
  return types.NewList()
}
