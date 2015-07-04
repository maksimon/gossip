package environment

import (
  "gossip/types"
)

var GlobalScope Scope

func init() {
  GlobalScope = Scope {
    map[string] types.LispElement{
      "+" : types.NewFunction(add),
      "-" : types.NewFunction(subtract),
      "*" : types.NewFunction(multiply),
      "/" : types.NewFunction(divide),
      "=" : types.NewFunction(equals),
      ">" : types.NewFunction(greaterThan),
      ">=": types.NewFunction(greaterThanEquals),
      "<" : types.NewFunction(lessThan),
      "<=": types.NewFunction(lessThanEquals),
    },
    nil,
  }
}
