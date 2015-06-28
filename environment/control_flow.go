package environment 

import (
  "gossip/types"
)

func If(input []types.LispElement) types.LispElement {
  if len(input) < 2 {
    panic("Error. Not enough arguments to 'if'");
  }
  if input[0].Type() != types.NilType {
    return input[1]
  }
  if len(input) > 1 {
    return input[2]
  }
  return types.NewNil()
}
