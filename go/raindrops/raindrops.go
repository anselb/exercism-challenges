package raindrops

import "strconv"

func Convert(input int) string {
  var expected string

  if input % 3 == 0 {
    expected += "Pling"
  }
  if input % 5 == 0 {
    expected += "Plang"
  }
  if input % 7 == 0 {
    expected += "Plong"
  }
  if expected == "" {
    expected += strconv.Itoa(input)
  }

  return expected
}
