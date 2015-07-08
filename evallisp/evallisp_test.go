package evallisp

import (
  "testing"
  "gossip/types"
  "gossip/environment"
  "gossip/parselisp"
  "fmt"
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
    "(+ 6"          +
      "(if (> 2 1)" +
        " 2"        +
        " 1000)"    +
      "(if (< 2 1)" + 
        " 1000"     +
        " 10)"      +
     ")"         
  parsed := parselisp.ParseLisp(program)
  program_value := eval(environment.GlobalScope, &parsed)
  if program_value.Label() != "18" {
    t.Error(fmt.Sprintf("If does not work as expected. Calculation was %s and not 18", program_value.Label()))
  }
}

func TestQuote(t * testing.T) {
  program := "(quote (1 + 3 (+ 2 4)))"
  parsed := parselisp.ParseLisp(program)
  evaled_program := eval(environment.GlobalScope, &parsed).(*types.LispList)
  if evaled_program.At(0).Label() != "1" {
    t.Error("quote did not return evalled function correctly")
  }
  if evaled_program.At(1).Label() != "(function)" {
    t.Error("quote did not return evalled function correctly")
  }
  if evaled_program.At(2).Label() != "3" {
    t.Error("quote did not return evalled function correctly")
  }
  if evaled_program.At(3).Label() != "6" {
    t.Error("quote did not return evalled function correctly")
  }
}

func TestEvalDefun(t * testing.T) {
  program :=
	  "(quote ("                  +
	    "(defun (pxthree (x y)"   +
	      "(* (+ x y) 3)))"       +
      "(defun (tryagain (x y)"  +
        "(+ (pxthree x y) 2)))" +
	    "(+ 4 (tryagain 3 4))"    +
	  "))"
  parsed := parselisp.ParseLisp(program)
  evaled_program := eval(environment.GlobalScope, &parsed).(*types.LispList)
  if evaled_program.At(2).Label() != "27" {
    t.Error(fmt.Sprintf("cannot defun. Got %s", evaled_program.At(2).Label()))
  }
}

func TestEvalSet(t *testing.T) {
  program :=
    "(quote ("        +
      "(set max 711)" +
      "(+ max 10)"    +
      "(set max 200)" +
      "(+ max 1)"     +
    "))"
  parsed := parselisp.ParseLisp(program)
  evaled_program := eval(environment.GlobalScope, &parsed).(*types.LispList)
  if evaled_program.At(1).Label() != "721" {
    t.Error("set does not work")
  }
  if evaled_program.At(3).Label() != "201" {
    t.Error("set does not work: cannot reset variable")
  }
}

func TestRecursiveDefun(t *testing.T) {
  program :=
    "(quote ("                                    +
      "(defun (fib (n sec fst)"                   +
        "(if (> n 0)"                             +
          "(fib (- n 1) (+ sec fst) sec)" +
          "sec"                                   +
        ")"                                       +
      "))"                                        +
      "(fib 4 1 1)"                               +
    "))"
  parsed := parselisp.ParseLisp(program)
  evaled_program := eval(environment.GlobalScope, &parsed).(*types.LispList)
  if evaled_program.At(1).Label() != "8" {
    t.Error(fmt.Sprintf("cannot recursivly defun. Got %s", evaled_program.At(1).Label()))
  }
}

func TestDefun(t * testing.T) {
  definition := parselisp.ParseLisp("(pxthree (x y) (* (+ x y) 3))")
  args:=  parselisp.ParseLisp("(5 2)")

  function_name, function:= defun(definition.Children, &environment.GlobalScope)

  result := function.Operate(args.Children)
  if result.Label() != "21" || function_name != "pxthree" {
    t.Error("function declarations don't work")
  }
}
