package evallisp

import (
  "gossip/types"
  "gossip/environment"
  "fmt"
)

func eval(scope environment.Scope, element types.LispElement) types.LispElement {
  ret := element
  if (element.Type() == types.RuneType) {
    var ok bool
    ret , ok = scope.LookupVariable(element.Label()) 
    if !ok {
      panic(fmt.Sprintf("Symbol not defined. Expecting variable, but found %s", ret.Label()))
    }
  } else if (element.Type() == types.ListType) {
    listElement := element.(*types.LispList)
    var ok bool
    if(listElement.Length() == 0 || listElement.At(0).Type() != types.RuneType) {
      return element
    }
    function, ok := scope.LookupFunction(listElement.At(0).Label())
    if !ok {
      panic(fmt.Sprintf("Symbol not defined. Expecting function, but found %s", ret.Label()))
    }
    args := types.NewList()
    for _ , val := range (element.(*types.LispList)).Children[1:] {
      args.Append(eval(scope,val))
    }
    ret = function(args.Children)
  }
  return ret
}
