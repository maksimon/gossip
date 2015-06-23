package types

import (
  "strconv"
)

type LispType int
const (
  NumberType LispType = iota
  FunctionType LispType = iota 
  ListType LispType = iota
  RuneType LispType = iota
)

type LispElement interface {
  Label() string
  Type() LispType
}

type LispPrimative struct {
  label string
  lispType LispType
}

func IsPrimative(inputType LispType) bool {
  if (inputType == NumberType) {
    return true;
  }
  return false
}

func NewRune(runeString string) * LispPrimative {
  return &LispPrimative{ runeString, RuneType }
}

func (primative *LispPrimative) Label() string {
  return primative.label
}

func (primative *LispPrimative) Type() LispType {
  return primative.lispType
}

type LispNumber struct {
  LispPrimative
  value float64
}

func (number *LispNumber) Value() float64 {
  return number.value;
}

func NewNumberFromLabel(numberString string) *LispNumber {
  numberValue, _ := strconv.ParseFloat(numberString, 64)
  return &LispNumber{
    LispPrimative{ numberString, NumberType },
    numberValue,
  }
}

func NewNumberFromValue(numberValue float64) *LispNumber {
  numberLabel := strconv.FormatFloat(numberValue, 'f', -1, 64)
  return &LispNumber {
    LispPrimative{numberLabel, NumberType},
    numberValue,
  }
}

type LispList struct {
  LispPrimative
  Children []LispElement
}

func (list *LispList) Append(element LispElement) {
  list.Children = append(list.Children, element)
}

func (list *LispList) At (index int) LispElement {
  return list.Children[index]
}

func (list *LispList) Length () int {
  return len(list.Children);
}

func NewList() *LispList {
  return &LispList { 
    LispPrimative { "(list)" , ListType },
    make([]LispElement, 0),
  }
}
