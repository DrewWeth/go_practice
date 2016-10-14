package main

import "fmt"
import "time"

func main(){
  fmt.Println("Welcome")
  var num int = 1
  name := "Drew"
  fmt.Println(name, num)

  for i := 0; i < 5; i++{
    if i % 2 == 0{
      fmt.Println(i, "Event")
    }else{
      fmt.Println(i, "Odd")
    }


  }

  t := time.Now()
  switch{
  case t.Hour() < 12:
    fmt.Println("Before noon")
  default:
    fmt.Println("After noon")

  }

  arr := [4]int{1,2,3}
  fmt.Println(arr)

  slice := make([]string, 2)
  fmt.Println(slice)
  // var s []string

  m := make(map[string]int)
  m["one"] = 1
  m["two"] = 2
  fmt.Println("m", m)

  a,b := m["three"]
  fmt.Println("a", a, "b", b)

  m2 := map[string]string{"drew": "cool"}
  fmt.Println(m2)

  nums := []int{1,2,3}
  sum := 0
  for _, num := range nums{
    sum += num
  }
  fmt.Println(sum)

  for k, v := range m{
    fmt.Println(k, "->", v)
  }

  for _, c := range "Hey"{
    fmt.Printf("%c\n", c)
  }

  seq := next()
  for a:=0; a < 3;a++{
    fmt.Println(seq())
  }

  r := rect{10,20}
  measure(r)

  go print("Hello world")
  go func(str string){
    fmt.Println(str)
  }("Test")


  // messages := make(chan string)


  c1 := make(chan string)
  c2 := make(chan string)

  go func(){
    time.Sleep(time.Second * 2)
    c1 <- "done 1"
  }()

  go func(){
    time.Sleep(time.Second * 1)
    c2 <- "done 2"
  }()

  for i:= 0; i < 2;i++{

    select {
    case msg := <- c1:
      fmt.Println(msg)
    case msg2 := <- c2:
      fmt.Println(msg2)
    }
  }


  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}

func add(a int, b int) int{
  return a+b
}

func next() func() int{
  i := 0
  return func() int{
    i++
    return i
  }
}

type geometry interface{
  area() float64
  perim() float64
}

type rect struct{
  width, height float64
}

func (r rect) area() float64{
  return r.width * r.height
}
func (r rect) perim() float64{
  return r.width*2  + r.height*2
}

func measure (g geometry){
  fmt.Println("Area", g.area())
  fmt.Println("Perim", g.perim())

}

func print (str string){
  for i := 0; i < len(str)-1; i++{
    fmt.Println(str, ":", i)
  }
}
