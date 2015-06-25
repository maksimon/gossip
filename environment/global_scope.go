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
    },
    map[string] types.LispElement {
    },
    nil,
  }
}
