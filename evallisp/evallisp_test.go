package evallisp

import (
  "testing"
  "gossip/types"
  "gossip/environment"
  "gossip/parselisp"
)

func TestEvalArithmetic(t * testing.T) {
  program := "(+ 1 2 (+ 0 (* 2 4) 1))"
  parsed := parselisp.ParseLisp(program)
  newElem := eval(environment.GlobalScope, &parsed)
  if (newElem.(*types.LispNumber)).Value() != 12 {
    t.Error("Arithmetic functions not evaluated correctly")
  }
}

func TestEvalIf(t * testing.T) {
  program := 
    "(+ 6"       +
      "(if 1"    +
        " 2"     +
        " 1000"  +
      ")"        +
      "(if nil"  + 
        " 1000"  +
        " 10)"   +
     ")"         
  parsed := parselisp.ParseLisp(program)
  program_value := eval(environment.GlobalScope, &parsed)
  if (program_value.(*types.LispNumber)).Value() != 18 {
    t.Error("If does not work as expected")
  }
}

func TestDefun(t * testing.T) {
  definition := parselisp.ParseLisp("(pxthree (x y) (* (+ x y) 3))")
  args:=  parselisp.ParseLisp("(5 2)")

  function:= defun(definition.Children, &environment.GlobalScope)

  result := function.Operate(args.Children)
  if result.Label() != "21" {
    t.Error("function declarations don't work")
  }
}
