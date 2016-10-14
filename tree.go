package main

import  "fmt"
// import "math"

type Node struct{
  left *Node
  right *Node
  value int
  height int
}

func main(){
  head := &Node{value:1, height:1}
  head.left = &Node{value:2}
  head.right = &Node{value:3}
  head.right.left = &Node{value:4}
  head.right.right = &Node{value:5}
  head.right.left.left = &Node{value:6}

  fmt.Println("Head", head)

  // bfs(head)
  // fmt.Println(height(head, 1))
  dfs(head)
}

func bfs(head *Node){
  queue := make([]*Node,0)
  queue = append(queue, head)

  for len(queue) > 0{
    node := queue[0]
    queue = queue[1:]
    fmt.Println("Height", node.height, node)

    if node.left != nil{
      node.left.height = node.height + 1
      queue = append(queue, node.left)
    }
    if node.right != nil{
      node.right.height = node.height + 1
      queue = append(queue, node.right)
    }
  }
}

func height(node *Node, level int) int{
  if node == nil{
    return 0
  }else{
    return 1 + Max(height(node.left, level),height(node.right, level))
  }
}


func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func dfs(node *Node){
  if node == nil{
      return
  }
  dfs(node.left)
  fmt.Println(node.value)
  dfs(node.right)
}
