package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  strings := strings.Split(s, " ")
  
  wordCount := make(map[string] int)
  
  for _, word := range strings {
    count := wordCount[word]
    wordCount[word] = count + 1
  }

  return wordCount
}

func main() {
  wc.Test(WordCount)
}
