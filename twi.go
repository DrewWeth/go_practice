// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"
import _ "testing"

var p = fmt.Println

type Node struct {
	left  *Node
	right *Node
	value string
}

func bfs(root *Node, target string) bool {
	if root == nil {
		return false
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if target == node.value {
			return true
		}
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return false
}

func main() {
	p(test("A", true, testBasicTree()))
	p(test("F", false, testBasicTree()))
	p(test("", true, testBasicTree()))
	p(test("A", false, nilTree()))
}

func test(input string, expected bool, root *Node) bool {
	if bfs(root, input) == expected {
		return true
	}
	return false
}

func testBasicTree() *Node {
	root := &Node{nil, nil, "A"}
	root.left = &Node{nil, nil, "B"}
	root.left.left = &Node{nil, nil, "D"}
	root.right = &Node{nil, nil, "C"}
	root.right = &Node{nil, nil, ""}
	return root
}

func nilTree() *Node {
	return nil
}
