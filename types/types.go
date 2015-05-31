package types

type LispType int

type LispElem struct {
  Value string
  Children *[]LispElem
  Type LispType
}

const (
  LispNumber LispType = iota
  LispFunction LispType = iota 
  LispArray  LispType = iota
)

