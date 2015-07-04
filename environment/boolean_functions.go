package environment

import (
 "gossip/types"
)

func equals(input []types.LispElement) types.LispElement {
  curValue := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    if curValue != (val.(*types.LispNumber)).Value() {
      return types.NewList()
    }
  }
  return types.NewNumberFromValue(1)
}

func greaterThan(input []types.LispElement) types.LispElement {
  curValue := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    if curValue <= (val.(*types.LispNumber)).Value() {
      return types.NewList()
    }
    curValue = (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(1)
}

func greaterThanEquals(input []types.LispElement) types.LispElement {
  curValue := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    if curValue < (val.(*types.LispNumber)).Value() {
      return types.NewList()
    }
    curValue = (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(1)
}

func lessThan(input []types.LispElement) types.LispElement {
  curValue := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    if curValue >= (val.(*types.LispNumber)).Value() {
      return types.NewList()
    }
    curValue = (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(1)
}

func lessThanEquals(input []types.LispElement) types.LispElement {
  curValue := (input[0].(*types.LispNumber)).Value()
  for _, val := range input[1:] {
    if (val.(*types.LispNumber)).Value() > curValue {
      return types.NewList()
    }
    curValue = (val.(*types.LispNumber)).Value()
  }
  return types.NewNumberFromValue(1)
}
