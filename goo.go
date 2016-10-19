package main
import "fmt"
import "strings"

var p = fmt.Println
var input = `
dir1
 dir11
 dir12
  file.txt
  pic.jpg
 dir3
  pic2.jpg
dir2
 dir3
  dir4
   dir5
    pic4.jpg
`

func main(){
    p(solution(input))
    p("Test passes:", test(solution(input), "/dir2/dir3/dir4/dir5"))
}

// Iterates through string directory. Each folder/file is separated by new lines with leading white space indicating directory depth. Returns longest path to file with .jpg extension
func solution(S string) string{
  maxLength := ""
  lines := strings.Split(S, "\n")
  stack := []string{}
  for _, e := range(lines){
    if strings.Index(e, ".") > 0{
      if strings.Index(e, ".jpg") > 0{
        stackStr := strings.Join(stack, "/")
        if len(stackStr) > len(maxLength){
          maxLength = stackStr
        }
      }
    }else{
      for len(stack) > countSpaces(e){
        stack = stack[0:len(stack)-1] // Pop off stack
      }
      if len(stack) == countSpaces(e){
        stack = append(stack, strings.Trim(e, " ")) // Add to stack
      }
      // Good for testing. Prints directory state during search
      // p(stack)
    }
  }
  return "/" + maxLength
}

func test(output, expected string) bool{
  return output == expected
}

func countSpaces(S string) int{
  counter := 0
  arr := strings.Split(S, "")
  for _, e := range arr{
    if e == " " {
      counter += 1
    }else{
      return counter
    }
  }
  return counter
}
