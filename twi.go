// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

var print = fmt.Println

type Node struct {
	left  *Node
	right *Node
	value string
}

// Performs breadth first search on a tree given a root node. Returns true iff string value exists in tree
func bfs(root *Node, target string) bool {
	if root == nil {
		return false
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		// Pop off queue element and store in node.
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
	print(test("A", true, testBasicTree()))
	print(test("F", false, testBasicTree()))
	print(test("", true, testBasicTree()))
	print(test("A", false, nil))
}


// Tests to ensure bfs returns expected value given a root node to a tree
func test(input string, expected bool, root *Node) bool {
	return bfs(root, input) == expected
}

// Creates basic tree and returns the root node
func testBasicTree() *Node {
	root := &Node{nil, nil, "A"}
	root.left = &Node{nil, nil, "B"}
	root.left.left = &Node{nil, nil, "D"}
	root.right = &Node{nil, nil, "C"}
	return root
}
