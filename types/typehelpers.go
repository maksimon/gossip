package types

func IsNil(element LispElement) bool {
  if element.Type() == ListType {
    elementLength := (element.(*LispList)).Length()
    return (elementLength == 0);
  }
  return false
}

func IsPrimative(inputType LispType) bool {
  if (inputType == NumberType) {
    return true;
  }
  return false
}
