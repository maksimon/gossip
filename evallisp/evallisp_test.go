package evallisp

import (
  "testing"
  "gossip/types"
  "gossip/environment"
  "gossip/parselisp"
)

func TestEvalPlus(t * testing.T) {
  program := "(+ 1 2 (+ 0 (* 2 4) 1))"
  parsed := parselisp.ParseLisp(program)
  newElem := eval(environment.GlobalScope, &parsed)
  if (newElem.(*types.LispNumber)).Value() != 12 {
    t.Error("Fail")
 }
}

