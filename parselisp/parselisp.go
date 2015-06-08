package parselisp 

import(
  "regexp"
  "gossip/types"
)

var tokenRegexp *regexp.Regexp = regexp.MustCompile(`(\(|\)|[^() ]+)`)

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
