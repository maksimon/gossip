package types

type LispNil struct {
  LispPrimative
}

func NewNil() *LispNil {
  return &LispNil{
    LispPrimative{"nil", NilType},
  }
}
