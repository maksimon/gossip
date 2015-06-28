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

func TestIf(t * testing.T) {
  program := 
  "(+"           +
    "("          +
      "if 1"   +
        " 4"    +
    ")"          +
    "2"        + 
    "("          +
      "if nil" + 
        " 1000" +
        " 10"   +
    ")"          +
  ")"
  parsed := parselisp.ParseLisp(program)
  program_value := eval(environment.GlobalScope, &parsed)
  //program_value := parsed.At(2)
  //if (program_value.Label() != "nil") {
  //  t.Error(fmt.Sprintf("Using poppin fresh as an adjective %s", program_value.Label()))
  //}
  if (program_value.(*types.LispNumber)).Value() != 16 {
    t.Error("If does not work as expected")
  }
}
