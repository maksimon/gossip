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

func TestEvalDefSet(t *testing.T) {
  program :=
    "(quote ("        +
      "(def dennis 711)" +
      "(+ dennis 10)"    +
      "(set dennis 200)" +
      "(+ dennis 1)"     +
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

func TestEvalLambda(t *testing.T) {
  program :=
    "(quote(" +
      "(def funcy (lambda (x y) (* (+ x y) 3)))" +
      "(+ 1 (funcy 4 5) )" +
   "))"
  parsed := parselisp.ParseLisp(program)
  evaled_program := eval(environment.GlobalScope, &parsed).(*types.LispList)
  if evaled_program.At(1).Label() != "28" {
    t.Error("set does not work")
  }
}

func TestRecursiveLambda(t *testing.T) {
  program :=
    "(quote ("                                    +
      "(def fib (lambda (n sec fst)"                   +
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

func TestLambda(t * testing.T) {
  function_args := parselisp.ParseLisp("(x y)")
  function_def  := parselisp.ParseLisp("(* (+ x y) 3)")

  custom_function := lambda(&function_args, &function_def,&environment.GlobalScope)
  result := custom_function.Operate(
    []types.LispElement{
      types.NewNumberFromValue(4),
      types.NewNumberFromValue(5),
    },
  )
  if result.Label() != "27" {
    t.Error("lambda does not work")
  }
}
