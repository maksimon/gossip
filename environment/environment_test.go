package environment

import (
  "testing"
  "gossip/types"
)

func TestAritmethic(t *testing.T) {
  success := true
  two := types.NewNumberFromLabel("6")
  three := types.NewNumberFromLabel("2")

  list := types.NewList() 
  list.Append(two)
  list.Append(three)

  sum        := GlobalScope.Functions["+"](list.Children).(*types.LispNumber)
  difference := GlobalScope.Functions["-"](list.Children).(*types.LispNumber)
  product    := GlobalScope.Functions["*"](list.Children).(*types.LispNumber)
  quotient   := GlobalScope.Functions["/"](list.Children).(*types.LispNumber)


  if !(sum.Label() == "8") && !(sum.Value() == 8) {
    success = false
  }
  if !(difference.Label() == "4") && !(difference.Value() == 4) {
    success = false
  }
  if !(product.Label() == "12") && !(product.Value() == 12) {
    success = false
  }
  if (quotient.Label() != "3") && (quotient.Value() != 3) {
    success = false
  }

  if !success {
    t.Error("Fail");
  }
}
