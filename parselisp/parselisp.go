package main

import(
  "fmt"
  "regexp"
  "gossip/types"
)

var tokenRegexp *regexp.Regexp = regexp.MustCompile(`(\(|\)|[^ ()]+)`)
var raw_source string = "(print (append your total is :(+ 1 2 3 (+ 4 5 6 8) (* 3 4) ) ) )"
//func ParseLisp(/*raw_source string*/) types.LispElement {
//  helper := func () types.LispElement {
//    currentElement := types.LispElement{ "", &[]types.LispElement{}, types.LispList }
//    return currentElement
//  }
//  return helper();
//}
//
//func main() {
//  fmt.Println(tokenRegexp.FindAllString(raw_source,-1)[1])
//  lt := types.LispElement{ "-", nil ,types.LispFunction}
//  fmt.Printf("The value is %d\n", lt.Type)
//  //return lt
//  //ParseLisp("(print (append your total is :(+ 1 2 3 (+ 4 5 6 8) (* 3 4) ) ) )")
//  le2 := ParseLisp()
//  fmt.Printf("The valie is also %d\n", le2.Type)
//}

func ParseLisp(rawSource string) types.LispList {
  ret := types.NewList()
  return *ret

}

func parseLispHelper(rawSource *string, index int) types.LispList, int {
  retList := types.NewList()
  for {
    switch *rawSource[index] {
      default:
        retList.Children = retList.Children.append(types.NewNumber(*rawSource[index]))
      case "(":
        child, index := parseLispHelper(rawSource, index)
        retList.Children = retList.Children.append(child)
      case ")":
        return retList, index
    }
    index += 1
  }
}

func main() {
  lp := *types.NewNumber("+")
  ll := ParseLisp("hey")
  fmt.Printf("The type is %d\n", lp.Type())
  fmt.Printf("The length is %d\n", len(ll.Children))
}
