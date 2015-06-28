package types

import (
  "testing"
)

func TestIsNil(t *testing.T) {
  number := NewNumberFromValue(2)

  nillist := NewList()

  notnillist := NewList()
  notnillist.Append(number)

  if IsNil(number) {
    t.Error("IsNil thinks that a number is nil")
  }
  if IsNil(notnillist) {
    t.Error("IsNil thinks that a non-empty list is nil")
  }
  if !(IsNil(nillist)) {
    t.Error("IsNil thinks that a nil list is not nil")
  }
}
