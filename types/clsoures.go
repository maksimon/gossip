package types

type LispFunction struct {
  LispPrimative
  operation func([]LispElement) LispElement
}

func (function *LispFunction) Operate(args []LispElement) LispElement {
  return function.operation(args)
}

func NewFunction(op func([]LispElement) LispElement) *LispFunction {
  return &LispFunction {
    LispPrimative  { "(function)", FunctionType },
    op,
  }
}
