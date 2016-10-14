package main

import "fmt"
var p = fmt.Println

func main(){
  fmt.Println("Starting")

  input := gather()
  fmt.Println(input)

  p(edits1(input))
}

func gather() string{
  var input string
  fmt.Scanln(&input)
  return input
}

func edits1(str string)[]string{
  alphabet := "abcdefghijklmnopzrstuvwxyz"
  deletions := func(str string)[]string{
    for i, _ := range alphabet{
      p(i, alphabet[0:i+1])
    }
    return []string{"a"}
  }(str)

  // p(alphabet[)
  return deletions
}
