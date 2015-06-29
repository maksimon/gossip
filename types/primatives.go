package types

type LispType int
const (
  NumberType LispType = iota
  ListType LispType = iota
  RuneType LispType = iota
  FunctionType LispType = iota
)

type LispElement interface {
  Label() string
  Type() LispType
}

type LispPrimative struct {
  label string
  lispType LispType
}

func (primative *LispPrimative) Label() string {
  return primative.label
}

func (primative *LispPrimative) Type() LispType {
  return primative.lispType
}

func NewRune(runeString string) * LispPrimative {
  return &LispPrimative{ runeString, RuneType }
}
