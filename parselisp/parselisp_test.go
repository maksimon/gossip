package parselisp

import (
  "testing"
  "gossip/types"
)

func verifyLispElementValues (l []types.LispElement, v []string) bool {
  for i := 0 ; i < len(v) ; i++ {
    if l[i].Value() != v[i] {
      return false
    }
  }
  return true
}



func TestParseNumberList(t *testing.T) {
  program := "(1 2 (8 9 10) 3)"
  parsed := ParseLisp(program)
  success := true

  success = (parsed.At(2).Type() == types.ListType )
  sublist := parsed.At(2).(*types.LispList)

  success = verifyLispElementValues(parsed.Children[0:2] , []string{"1", "2"})
  success = verifyLispElementValues(sublist.Children, []string{"8", "9", "10"})
  success = verifyLispElementValues(parsed.Children[3:], []string{"3"})

  if !success {
    t.Error("Fail")
  }
}
