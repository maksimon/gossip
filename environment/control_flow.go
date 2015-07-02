package environment 

import (
  "gossip/types"
)

func lispIf(input []types.LispElement) types.LispElement {
  if len(input) < 2 {
    panic("Error. Not enough arguments to 'if'");
  }
  if !(types.IsNil(input[0])) {
    return input[1]
  }
  if len(input) > 1 {
    return input[2]
  }
  return types.NewList()
}

var If *types.LispFunction

func init() {
  If = types.NewFunction(lispIf)
}
