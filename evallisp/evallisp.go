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
    listElement := element.(*types.LispList)
    if listElement.Length() == 0 {
      return element;
    }
    list_leader := listElement.At(0).Label()
    if (list_leader == "quote") {
      return quote(listElement.At(1), scope)
    } else if (list_leader == "if") {
      return If(listElement.Children[1:], &scope)
    } else if (list_leader == "def" || list_leader == "set") {
      if (listElement.At(1).Type() == types.RuneType && listElement.Length() >= 2) {
        var_name  := listElement.At(1).Label()
        var_value := eval(scope, listElement.At(2))
        if list_leader == "def" {
          scope.AddVariable(var_name,var_value)
        }
        if list_leader == "set" {
          scope.SetVariable(var_name,var_value)
        }
      } else {
        panic("Improperly formatted set")
      }
    } else if (list_leader == "defun") {
      defun_spec := listElement.At(1).(*types.LispList)
      function_name, function_def := defun(defun_spec.Children, &scope)
      scope.AddVariable(function_name,function_def)
      return function_def
    } else { //function stuff 
	    var ok bool
	    function, ok := scope.LookupFunction(list_leader)
	    if !ok {
	      panic(fmt.Sprintf("Symbol not defined. Expecting function, but found %s", ret.Label()))
	    }
	    return function.Operate(evalArgs(listElement.Children[1:]));
    }
  }
  return ret
}

