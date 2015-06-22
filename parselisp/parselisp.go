package parselisp 

import(
  "regexp"
  "gossip/types"
)

var tokenRegexp *regexp.Regexp = regexp.MustCompile(`(\(|\)|[^() ]+)`)
var numberRegexp *regexp.Regexp = regexp.MustCompile(`[0-9]+`);

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
    switch {
      default:
        retList.Append(types.NewRune(cursorValue))
      case numberRegexp.MatchString(cursorValue):
         retList.Append(types.NewNumberFromLabel(cursorValue))
      case cursorValue == "(":
        var child *types.LispList
        child, cursor = parseLispHelper(rawSource, cursor)
        retList.Append(child)
      case cursorValue == ")":
        return retList, cursor
    }
    cursor += 1
  }
}
