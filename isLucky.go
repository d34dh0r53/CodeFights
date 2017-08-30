package main

import (
  "fmt"
  "strconv"
)

func isLucky(n int) bool {
  ticket := strconv.Itoa(n)
  var left int = 0
  var right int = 0

  for i := 0; i < len(ticket) / 2; i++ {
    //left += strconv.Atoi(string(toString[i]))
    lString := string(ticket[i])
    lInt, _ := strconv.Atoi(lString)
    left += lInt
  }

  for j := len(ticket) / 2; j < len(ticket); j++ {
    rString := string(ticket[j])
    rInt, _ := strconv.Atoi(rString)
    right += rInt
  }

  if left == right {
    return true
  } else {
    return false
  }
}

func main() {
  var ticketNumber int = 293017
  fmt.Println(isLucky(ticketNumber))
}
