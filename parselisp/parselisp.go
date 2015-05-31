package main

import(
  "fmt"
  "regexp"
  "gossip/types"
)

/*

func makeRootNode syntaxNode {

}


func countTokens(raw_source string) {


}*/

func main() {
  re := regexp.MustCompile(`(\(|\)|[^ ()]+)`)
  fmt.Println(re.FindAllString("(print (append your total is :(+ 1 2 3 (+ 4 5 6 8) (* 3 4) ) ) )",-1)[3])
  lt := types.LispElem{ "-", nil ,types.LispFunction}
  fmt.Printf("The value is %d\n", lt.Type)

}

