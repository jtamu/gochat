package lib

import (
  "strconv"
)

func ToS(i uint) (s string) {
  s = strconv.Itoa(int(i))
  return
}

func ToI(s string) (i uint) {
  j, _ := strconv.Atoi(s)
  return uint(j)
}
