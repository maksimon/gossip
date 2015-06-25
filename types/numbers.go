package types

import (
  "strconv"
)

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
