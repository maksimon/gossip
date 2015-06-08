package main

import(
  "fmt"
  "regexp"
  "gossip/types"
)

var tokenRegexp *regexp.Regexp = regexp.MustCompile(`(\(|\)|[^() ]+)`)
//var raw_source string = "(print (append your total is :(+ 1 2 3 (+ 4 5 6 8) (* 3 4) ) ) )"

func ParseLisp(rawSource string) types.LispList {
  parsedSource := tokenRegexp.FindAllString(rawSource, -1)
  ret , _ := parseLispHelper((&parsedSource), 0)
  return *ret
}

func parseLispHelper(rawSource *[]string, cursor int) (*types.LispList, int) {
  retList := types.NewList()
  cursor += 1 // skipping over first '('
  var cursorValue string
  for {
    cursorValue = string((*rawSource)[cursor])
    switch cursorValue {
      default:
         retList.Append(types.NewNumber(cursorValue))
      case "(":
        var child *types.LispList
        child, cursor = parseLispHelper(rawSource, cursor)
        retList.Append(child)
      case ")":
        return retList, cursor
    }
    cursor += 1
  }
}

func main() {
  lp := *types.NewNumber("+")
  ll := ParseLisp("(1 2 (8 9 10) 3)")
  fmt.Printf("The type is %d\n", lp.Type())
  fmt.Printf("The length is %d\n", len(ll.Children))
  for i := 0; i < len(ll.Children) ; i++ {
    fmt.Printf("Value: %s Type: %d", ll.Children[i].Value(), ll.Children[i].Type())
    if (ll.Children[i].Type() == types.ListType) {
      listChild := ll.Children[i].(types.LispList)
      fmt.Printf(" Length: %d", len(listChild.Children))
    }
    fmt.Printf("\n")
  }
}
