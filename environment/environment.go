package environment

import (
  "gossip/types"
)

type Scope struct {
  Functions map[string] func([]types.LispElement) *types.LispElement
  Variables map[string] *types.LispElement
}

GlobalScope := Scope {
  map[string] func([]types.LispElement) *types.LispElement{
    "+" : Add,
    "-" : Subtract,
    "*" : Multiply,
    "/" : Divide,
  }
  map[string] *types.LispElement {
  }
}
