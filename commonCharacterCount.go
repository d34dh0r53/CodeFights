package main

import (
  "fmt"
)

func commonCharacterCount(s1 string, s2 string) int {

  type Key struct {
    Char byte
  }

  var commonCharacters int = 0

  s1CharacterCount := make(map[Key]int)
  s2CharacterCount := make(map[Key]int)

  for i := 0; i < len(s1); i++ {
    charCode := s1[i]
    s1CharacterCount[Key{charCode}]++
  }

  for i := 0; i < len(s2); i++ {
    charCode := s2[i]
    s2CharacterCount[Key{charCode}]++
  }

  for p := range s1CharacterCount {
    if s2CharacterCount[p] > 0 {
      if s1CharacterCount[p] < s2CharacterCount[p] {
        commonCharacters += s1CharacterCount[p]
      } else {
        commonCharacters += s2CharacterCount[p]
      }
    }
  }

  return commonCharacters
}

func main() {
  const exampleString = "aabcc"
  const exampleString2 = "adcaa"

  fmt.Println(commonCharacterCount(exampleString, exampleString2))
}
