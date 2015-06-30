package evallisp

import (
  "gossip/types"
  "gossip/environment"
  "fmt"
)

func defun(func_spec []types.LispElement, parent_scope *environment.Scope) *types.LispFunction {
  function_name := func_spec[0].(*types.LispPrimative)
  function_arguments := func_spec[1].(*types.LispList)
  function_contents := func_spec[2].(*types.LispList)

  if !validFunctionName(function_name) {
    panic(fmt.Sprintf("Improper function definition. Cannot use %s as function name", function_name.Label()))
  }

  for _, function_argument := range(function_arguments.Children) {
    if function_argument.Type() != types.RuneType {
      panic("Improper function definition")
    }
  }
  function := func(args []types.LispElement) types.LispElement {
    function_scope := environment.Scope {
      map[string] func([]types.LispElement) types.LispElement{},
      map[string] types.LispElement{},
      parent_scope,
    }
    for index, function_variable := range(function_arguments.Children) {
      function_scope.AddVariable(function_variable.Label(), args[index])
    }
    return eval(function_scope, function_contents)
  }

  return types.NewFunction(function)
}

func validFunctionName(name_element types.LispElement) bool {
  return (name_element.Type() == types.RuneType )
}
