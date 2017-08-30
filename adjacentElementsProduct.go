package main

import "fmt"

func adjacentElementsProduct(inputArray []int) int {
  n := len(inputArray)
  max := 0
  for count := 0; count < n-1; count++ {
    x := inputArray[count]
    y := inputArray[count+1]
    product := x * y
    fmt.Printf("X: %d\tY: %d\tProduct: %d\tMax: %d\n", x, y, product, max)
    if max < product {
      max = product
    }
  }
  return max
}

func main() {
  i := []int{3,6,-2,-5,7,3}
  fmt.Println(adjacentElementsProduct(i))
}
