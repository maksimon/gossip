package types

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

func (list *LispList) Length() int {
  return len(list.Children);
}

func NewList() *LispList {
  return &LispList { 
    LispPrimative { "(list)" , ListType },
    make([]LispElement, 0),
  }
}

func NewListFromElements(elements []LispElement) *LispList {
  return &LispList {
    LispPrimative { "(list)", ListType },
    elements,
  }
}

