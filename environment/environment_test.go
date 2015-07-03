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

  add, _      := GlobalScope.LookupFunction("+")
  subtract, _ := GlobalScope.LookupFunction("-")
  multiply, _ := GlobalScope.LookupFunction("*")
  divide, _   := GlobalScope.LookupFunction("/")

  sum        := add.Operate(list.Children).(*types.LispNumber)
  difference := subtract.Operate(list.Children).(*types.LispNumber)
  product    := multiply.Operate(list.Children).(*types.LispNumber)
  quotient   := divide.Operate(list.Children).(*types.LispNumber)


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
    map[string] types.LispElement {
      "age" : types.NewNumberFromValue(22),
    },
    &GlobalScope,
  }
  GlobalScope.AddVariable("year", types.NewNumberFromValue(2015))
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
    map[string] types.LispElement {
      "-" : Subtract,
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

  difference := subtract.Operate(list.Children)
  product := multiply.Operate(list.Children) 

  if !(difference.Label() == "4") {
    t.Error("local scope function lookup failed")
  }
  if !(product.Label() == "12") {
    t.Error("global scope function lookup failed")
  }
}
