package environment

import (
  "gossip/types"
)

func Add(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum += (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}

func Subtract(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum -= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
func Multiply(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum *= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
func Divide(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum /= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
