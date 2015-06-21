package environment

import (
  "gossip/types"
  "strconv"
)


func Add(input []types.LispElement) *types.LispElement {
  sum := (input[0].(types.LispNumber)).Value()
  for _, val := range input[1:] {
    val.(types.LispNumber)
    sum += val.Value()
  }
  sumString := strconv(sum, 'f', 64)
  return &types.LispNumber {
    LispPrimative{sumString, types.NumberType},
    sum,
  }
}

func Subtract(input []types.LispElement) *types.LispElement {
  difference := (input[0].(types.LispNumber)).Value()
  for _, val := range input[1:] {
    val.(types.LispNumber)
    difference -= (val.(types.LispNumber)).Value()
  }
  differenceString := strconv(sum, 'f', 64)
  return &types.LispNumber {
    LispPrimative{differenceString, types.NumberType},
    difference,
  }
}

func Multiply(input []types.LispElement) *types.LispElement {
  product := (input[0].(types.LispNumber)).Value()
  for _, val := range input[1:] {
    val.(types.LispNumber)
    product *= val.Value()
  }
  productString := strconv(sum, 'f', 64)
  return &types.LispNumber {
    LispPrimative{productString, types.NumberType},
    product,
  }
}

func Divide(input []types.LispElement) *types.LispElement {
  product := (input[0].(types.LispNumber)).Value()
  for _, val := range input[1:] {
    val.(types.LispNumber)
    product /= val.Value()
  }
  productString := strconv(sum, 'f', 64)
  return &types.LispNumber {
    LispPrimative{productString, types.NumberType},
    product,
  }
}
