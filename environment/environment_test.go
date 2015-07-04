package environment

import (
  "testing"
  "gossip/types"
  "fmt"
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

func TestBoolean(t *testing.T) {
  equal, _      := GlobalScope.LookupFunction("=")
  greater_than, _ := GlobalScope.LookupFunction(">")
  greater_than_equal, _ := GlobalScope.LookupFunction(">=")
  less_than, _   := GlobalScope.LookupFunction("<")
  less_than_equal, _   := GlobalScope.LookupFunction("<=")

  test_bools_on_list := func(list_to_test []types.LispElement, expected_results []bool) (bool, int) {
    test_results := []bool{
	     (equal.Operate(list_to_test).Label() == "1"),
	     (greater_than.Operate(list_to_test).Label() == "1"),
	     (greater_than_equal.Operate(list_to_test).Label() == "1"),
	     (less_than.Operate(list_to_test).Label() == "1"),
	     (less_than_equal.Operate(list_to_test).Label() == "1"),
    }
    for i := range expected_results {
      if test_results[i] != expected_results[i] {
        return false, i
      }
    }
    return true, -1 
  }
  var test_result bool
  var failure_index int 

  test_result, failure_index = test_bools_on_list([]types.LispElement{
    types.NewNumberFromValue(2),
    types.NewNumberFromValue(2),
    types.NewNumberFromValue(2),
  }, 
  []bool{true,false,true,false,true})
  if !test_result {
    t.Error(fmt.Sprintf("the %d(st/nd/rd/th) bool function tested didnt work on equal numbers", failure_index))
  }

  test_result, failure_index = test_bools_on_list([]types.LispElement{
    types.NewNumberFromValue(2),
    types.NewNumberFromValue(5),
    types.NewNumberFromValue(9),
  }, 
  []bool{false,true,true,false,false})
  if !test_result {
    t.Error(fmt.Sprintf("the %d(st/nd/rd/th) bool function tested didnt work on strictly increasing numbers", failure_index))
  }

  test_result, failure_index = test_bools_on_list([]types.LispElement{
    types.NewNumberFromValue(2),
    types.NewNumberFromValue(5),
    types.NewNumberFromValue(5),
  }, 
  []bool{false,false,true,false,false})
  if !test_result {
    t.Error(fmt.Sprintf("the %d(st/nd/rd/th) bool function tested didnt work on increasing numbers", failure_index))
  }

  test_result, failure_index = test_bools_on_list([]types.LispElement{
    types.NewNumberFromValue(9),
    types.NewNumberFromValue(5),
    types.NewNumberFromValue(2),
  }, 
  []bool{false,false,false,true,true})
  if !test_result {
    t.Error(fmt.Sprintf("the %d(st/nd/rd/th) bool function tested didnt work on strictly decreasing numbers", failure_index))
  }

  test_result, failure_index = test_bools_on_list([]types.LispElement{
    types.NewNumberFromValue(9),
    types.NewNumberFromValue(5),
    types.NewNumberFromValue(5),
  }, 
  []bool{false,false,false,false,true})
  if !test_result {
    t.Error(fmt.Sprintf("the %d(st/nd/rd/th) bool function tested didnt work on decreasing numbers", failure_index))
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
      "-" : types.NewFunction(subtract),
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
