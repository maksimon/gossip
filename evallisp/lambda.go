package evallisp

import (
  "gossip/types"
  "gossip/environment"
)

func lambda( function_arguments *types.LispList, function_contents *types.LispList, parent_scope *environment.Scope) *types.LispFunction {
  for _, function_argument := range(function_arguments.Children) {
    if function_argument.Type() != types.RuneType {
      panic("Improper function definition")
    }
  }
  function := func(args []types.LispElement) types.LispElement {
    function_scope := environment.Scope {
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
