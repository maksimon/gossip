package evallisp

import(
  "gossip/types"
  "gossip/environment"
)

func eval(scope environment.Scope, element types.LispElement) types.LispElement {
  ret := element
  if (element.Type() == types.RuneType) {
    return scope.Variables[element.Label()]
  }
  if (element.Type() == types.ListType) {
    listElement := element.(*types.LispList)
    if(listElement.At(0).Type() != types.RuneType) {
      return element
    }
    function := scope.Functions[listElement.At(0).Label()]
    args := types.NewList()
    for _ , val := range (element.(*types.LispList)).Children[1:] {
      args.Append(eval(scope,val))
    }
    return function(args.Children)
  }
  return ret
}
