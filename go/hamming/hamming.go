package hamming

import "errors"

func Distance(a, b string) (int, error) {
  differencesNum := 0

  if len(a) != len(b) {
    return differencesNum, errors.New("Sequence lengths are unequal")
  } else {
    for pos := range a {
      if a[pos] != b[pos] {
        differencesNum += 1
      }
    }

    return differencesNum, nil
  }
}
