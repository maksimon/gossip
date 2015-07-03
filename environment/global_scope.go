package environment

import (
  "gossip/types"
)

var GlobalScope Scope

func init() {
  GlobalScope = Scope {
    map[string] types.LispElement{
      "+" : Add,
      "-" : Subtract,
      "*" : Multiply,
      "/" : Divide,
    },
    nil,
  }
}
