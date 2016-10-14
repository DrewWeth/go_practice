package main

import "fmt"

var p = fmt.Println

func main(){
  p(1, isEven(1))
  p(2, isEven(2))
}

func isEven(n int) bool{
  if n & 1 == 0{
    return true
  }else{
    return false
  }
}
