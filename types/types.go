package types

type LispType int

const (
  NumberType LispType = iota
  FunctionType LispType = iota 
  ListType  LispType = iota
)

type LispElement interface {
  Value() string
  Type() LispType
}

type LispPrimative struct {
  value string
  lispType LispType
}

func (LispPrimative *LispPrimative) Value() string {
  return LispPrimative.value
}

func (LispPrimative *LispPrimative) Type() LispType {
  return LispPrimative.lispType
}

func NewFunction(functionName string) *LispPrimative {
  return &LispPrimative{functionName, FunctionType}
}

func NewNumber(numberString string) *LispPrimative {
  return &LispPrimative{numberString, NumberType}
}

type LispList struct {
  Children []LispElement
}

func (list LispList) Value() string {
  return "(list)"
}

func (list LispList) Type() LispType {
  return ListType
}

func (list *LispList) Append(element LispElement) {
  list.Children = append(list.Children, element)
 
}

func (list *LispList) At (index int) LispElement {
  return list.Children[index]
}

func NewList() *LispList {
  return &LispList{ make([]LispElement, 0) }
}

