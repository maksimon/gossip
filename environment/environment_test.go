package environment

import (
  "testing"
  "gossip/types"
)

func TestAritmethic(t *testing.T) {
  six := types.NewNumberFromLabel("6")
  two := types.NewNumberFromLabel("2")

  list := types.NewList() 
  list.Append(six)
  list.Append(two)

  sum        := GlobalScope.Functions["+"](list.Children).(*types.LispNumber)
  difference := GlobalScope.Functions["-"](list.Children).(*types.LispNumber)
  product    := GlobalScope.Functions["*"](list.Children).(*types.LispNumber)
  quotient   := GlobalScope.Functions["/"](list.Children).(*types.LispNumber)


  if !(sum.Label() == "8") || !(sum.Value() == 8) {
    t.Error("'+' arithmetic function failed")
  }
  if !(difference.Label() == "4") || !(difference.Value() == 4) {
    t.Error("'-' arithmetic function failed")
  }
  if !(product.Label() == "12") || !(product.Value() == 12) {
    t.Error("'*' arithmetic function failed")
  }
  if !(quotient.Label() == "3") || !(quotient.Value() == 3) {
    t.Error("'\\' arithmetic function failed")
  }
}

func TestVariableLookup(t *testing.T) {
  lookup_scope := Scope {
    map[string] func([]types.LispElement) types.LispElement{
    },
    map[string] types.LispElement {
      "age" : types.NewNumberFromValue(22),
    },
    &GlobalScope,
  }
  GlobalScope.Variables["year"] = types.NewNumberFromValue(2015)
  age, _ := lookup_scope.LookupVariable("age")
  year,_ := lookup_scope.LookupVariable("year")
  if age.Label() != "22" {
    t.Error("local scope variable lookup failed")
  }
  if year.Label() != "2015" {
    t.Error("global scope variable lookup failed")
  }
}

func TestFunctionLookup(t * testing.T) {
  lookup_scope := Scope {
    map[string] func([]types.LispElement) types.LispElement{
      "-" : Subtract,
    },
    map[string] types.LispElement {
    },
    &GlobalScope,
  }
  six := types.NewNumberFromLabel("6")
  two := types.NewNumberFromLabel("2")
  list := types.NewList() 
  list.Append(six)
  list.Append(two)

  subtract, _ := lookup_scope.LookupFunction("-")
  multiply, _ := lookup_scope.LookupFunction("*")

  difference := subtract(list.Children)
  product := multiply(list.Children) 

  if !(difference.Label() == "4") {
    t.Error("local scope function lookup failed")
  }
  if !(product.Label() == "12") {
    t.Error("global scope function lookup failed")
  }
}
