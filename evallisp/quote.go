package evallisp

import (
  "gossip/types"
  "gossip/environment"
)

func quote(to_quote types.LispElement, scope environment.Scope) *types.LispList {
  to_quote_list := to_quote.(*types.LispList)
  quoted_list := types.NewList()
  for _, val := range(to_quote_list.Children) {
    quoted_list.Append(eval(scope,val))
  }
  return quoted_list
}
