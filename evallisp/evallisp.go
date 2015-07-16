package evallisp

import (
  "gossip/types"
  "gossip/environment"
  "fmt"
)

func eval(scope environment.Scope, element types.LispElement) types.LispElement {
	evalArgs := func(args []types.LispElement) []types.LispElement {
	  eval_args := make([]types.LispElement, 0)
	  for _, val := range(args) {
      eval_args = append(eval_args,eval(scope,val))
	  }
	  return eval_args
	}
  ret := element
  if (element.Type() == types.RuneType) {
    var ok bool
    ret , ok = scope.LookupVariable(element.Label()) 
    if !ok {
      panic(fmt.Sprintf("Symbol not defined. Expecting variable, but found %s", ret.Label()))
    }
  } else if (element.Type() == types.ListType) {
    list_element := element.(*types.LispList)
    if list_element.Length() == 0 {
      return element;
    }
    leader_label := list_element.At(0).Label()
    if (leader_label == "quote") {
      return quote(list_element.At(1), scope)
    } else if (leader_label == "if") {
      return If(list_element.Children[1:], &scope)
    } else if (leader_label == "def" || leader_label == "set") {
      if (list_element.At(1).Type() == types.RuneType && list_element.Length() >= 2) {
        var_name  := list_element.At(1).Label()
        var_value := eval(scope, list_element.At(2))
        if leader_label == "def" {
          scope.AddVariable(var_name,var_value)
        }
        if leader_label == "set" {
          scope.SetVariable(var_name,var_value)
        }
      } else {
        panic(fmt.Sprintf("Improperly formatted set. Label was (%s)", leader_label))
      }
    } else if (leader_label == "lambda") {
	    function_arguments := list_element.At(1).(*types.LispList)
	    function_contents  := list_element.At(2).(*types.LispList)
	    function_def := lambda(
        function_arguments,
        function_contents,
        &scope,
      )
      return function_def
    } else { //function stuff 
	    var ok bool
	    function, ok := scope.LookupFunction(leader_label)
	    if !ok {
	      panic(fmt.Sprintf("Symbol not defined. Expecting function, but found %s", ret.Label()))
	    }
	    return function.Operate(evalArgs(list_element.Children[1:]));
    }
  }
  return ret
}

