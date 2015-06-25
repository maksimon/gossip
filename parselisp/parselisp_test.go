package parselisp

import (
  "testing"
  "gossip/types"
)

func verifyLispElementLabels (l []types.LispElement, v []string) bool {
  for i := 0 ; i < len(v) ; i++ {
    if l[i].Label() != v[i] {
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

  success = verifyLispElementLabels(parsed.Children[0:2] , []string{"1", "2"})
  success = verifyLispElementLabels(sublist.Children, []string{"8", "9", "10"})
  success = verifyLispElementLabels(parsed.Children[3:], []string{"3"})

  if !success {
    t.Error("cannot parse numbers")
  }
}
