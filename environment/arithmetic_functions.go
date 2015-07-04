package environment

import (
  "gossip/types"
)

func add(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum += (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
func subtract(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum -= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
func multiply(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum *= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
func divide(input []types.LispElement) types.LispElement {
  sum := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    sum /= (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(sum)
}
