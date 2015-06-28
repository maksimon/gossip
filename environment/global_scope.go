package environment

import (
  "gossip/types"
)

var GlobalScope Scope

func init() {
  GlobalScope = Scope {
    map[string] func([]types.LispElement) types.LispElement{
      "+" : Add,
      "-" : Subtract,
      "*" : Multiply,
      "/" : Divide,
      "if": If,
    },
    map[string] types.LispElement {
    },
    nil,
  }
}
