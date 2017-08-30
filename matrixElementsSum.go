package main

import (
  "fmt"
)

func matrixElementsSum(matrix [][]int) int {
  rows := len(matrix)
  cols := len(matrix[0])
  for row := 0; row < rows; row++ {
    fmt.Printf("Row: %d Col: %v\n",row, matrix[row])
  }
  return cols
}

func main() {
  testMatrix := [][]int {{0,1,1,2},{0,5,0,0},{2,0,3,3}}
  fmt.Println(matrixElementsSum(testMatrix))
}
