package main
import s "strconv"
import "math/rand"
import "fmt"
import "time"
var p = fmt.Println

func main(){
  rand.Seed(time.Now().UnixNano())

  freq := make(map[string]int)

  for i := 0; i < 10;i++{
    arr := []int{1,2,3,4,5}
    count(shuffle(arr), freq)
  }
  p(freq)
}

func _swap(x, y int) (int, int){
  x = x ^ y
  y = x ^ y
  x = x ^ y
  return x,y
}

func shuffle(arr []int) []int{
  for i:=len(arr) - 1; i >= 0; i--{

    j:= rand.Intn(i+1) // between 0 (inclusive) and i (exclusive)
    arr[i], arr[j] = _swap(arr[i], arr[j])
  }

  return arr
}

func count(arr []int, hash map[string]int){
  for i, e := range arr{
    str := s.Itoa(i) + "-" + s.Itoa(e)
    p(str)
    hash[str] += 1
  }

}
