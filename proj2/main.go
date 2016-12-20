package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	conv "strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Car struct {
	id         int
	risk       float64
	valueLoss  float64
	horsepower float64
	cityMPG    float64
	highwayMPG float64
	price      float64
}

type Membership struct {
	name string
	attr [][]float64
}

type Node struct {
	decision string
	children []Node
	attr     string
	mem      int
}

var filename = "car_data.csv"
var dump = spew.Dump
var print = fmt.Println

// Reads input, setsup layers, forward pass and back-propogates for each test input. Finally, it prints the result
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cars := readFile()
	memberships := setupMemberships()
	root := setupTree()
	print("Root", root)
	// print(cars)
	// print(cars[i])
	// print(memberships)
	res := []float64{}
	for i, _ := range cars {
		fuzzySets := getFuzzySets(memberships, cars[i])
		print("Fuzzy set", fuzzySets)
		res = append(res, scanNode(root, fuzzySets))

	}
	// for i, r := range res {
	// 	print(r)
	// }
	m, i := max(res)
	print("max:", m, "car:", cars[i])
	print(getFuzzySets(memberships, cars[i]))

}

func getFuzzySets(memberships map[string]Membership, car Car) map[string][]float64 {
	fuzzySet := make(map[string][]float64)
	fuzzySet["risk"] = calculateMembership(memberships["risk"], car.risk)
	fuzzySet["valueLoss"] = calculateMembership(memberships["valueLoss"], car.valueLoss)
	fuzzySet["horsepower"] = calculateMembership(memberships["horsepower"], car.horsepower)
	fuzzySet["cityMPG"] = calculateMembership(memberships["cityMPG"], car.cityMPG)
	fuzzySet["highwayMPG"] = calculateMembership(memberships["highwayMPG"], car.highwayMPG)
	fuzzySet["price"] = calculateMembership(memberships["price"], car.price)
	return fuzzySet
}

func min(v []float64) (float64, int) {
	m := math.MaxFloat64
	index := -1
	for i, e := range v {
		if e < m {
			m = e
			index = i
		}
	}
	return m, index
}

func max(v []float64) (float64, int) {
	m := math.SmallestNonzeroFloat64
	index := -1
	for i, e := range v {
		if e > m {
			m = e
			index = i
		}
	}
	return m, index
}

// Not good name...
func scanChildren(children []Node, sets map[string][]float64) []float64 {
	res := []float64{}
	for i, _ := range children {
		res = append(res, scanNode(children[i], sets))
	}
	return res
}

func scanNode(node Node, sets map[string][]float64) float64 {
	if node.decision != "" {
		print(node.decision)
		if node.decision == "and" {
			min, _ := min(scanChildren(node.children, sets))
			return min
		} else if node.decision == "or" {
			max, _ := max(scanChildren(node.children, sets))
			return max
		} else {
			// Negation
			if len(node.children) != 1 {
				print("Problem 458.1")
			}
			return 1.0 - scanNode(node.children[0], sets)
		}
	} else {
		print(node.attr, "val:", sets[node.attr][node.mem])
		return sets[node.attr][node.mem]
	}
}

func setupTree() Node {
	// Left of root
	not := Node{decision: "not", children: []Node{Node{attr: "cityMPG", mem: 0}}}
	highway := Node{attr: "highwayMPG", mem: 2}
	horsepower := Node{attr: "horsepower", mem: 1}
	left := Node{decision: "and", children: []Node{not, highway, horsepower}}

	// Right of root
	risk := Node{attr: "risk", mem: 0}
	value := Node{attr: "valueLoss", mem: 0}
	price := Node{attr: "price", mem: 0}
	and1 := Node{decision: "and", children: []Node{risk, value}}
	right := Node{decision: "or", children: []Node{and1, price}}
	root := Node{decision: "and", children: []Node{left, right}}
	return root
}

func calculateMembership(membership Membership, value float64) []float64 {
	belong := []float64{}
	for _, attrs := range membership.attr {
		// print(attrs)
		if len(attrs) == 3 {
			belong = append(belong, triangle(attrs, value))
		} else if len(attrs) == 4 {
			belong = append(belong, trapezoid(attrs, value))
		} else {
			print("Problem 12.1")
		}
	}
	return belong
}

func trapezoid(attr []float64, value float64) float64 {
	if value <= attr[0] || value >= attr[3] {
		return 0.0
	} else {
		// In the trapezoid somewhere
		if value >= attr[1] && value <= attr[2] {
			// In the middle of trap
			return 1.0
		} else if value < attr[1] {
			// Left triangle part
			dist := attr[1] - attr[0]
			x := value - attr[0]
			return float64(x / dist)
		} else {
			// Right triangle part
			dist := attr[3] - attr[2]
			x := value - attr[2]
			return float64(x / dist)
		}
	}
}

func triangle(attr []float64, value float64) float64 {
	if value <= attr[0] || value >= attr[2] {
		return 0.0
	} else {
		// In the triangle somewhere
		if value < attr[1] {
			// Left side of 1
			dist := attr[1] - attr[0]
			x := value - attr[0]
			return float64(x / dist)
		} else {
			// Right side of 1
			dist := attr[2] - attr[1]
			x := attr[2] - value
			return float64(x / dist)
		}
	}
}

func setupMemberships() map[string]Membership {
	memberships := make(map[string]Membership)
	memberships["risk"] = Membership{"risk", [][]float64{{-3, -3, -2, 0}, {-2, 0, 2}, {0, 2, 3, 3}}}
	memberships["valueLoss"] = Membership{"valueLoss", [][]float64{{0, 0, 100, 120}, {100, 120, 200}, {120, 200, 300, 300}}}
	memberships["horsepower"] = Membership{"horsepower", [][]float64{{0, 0, 60, 100}, {60, 100, 140}, {100, 140, 250, 250}}}
	memberships["cityMPG"] = Membership{"cityMPG", [][]float64{{0, 0, 20, 30}, {20, 30, 40}, {30, 40, 60, 60}}}
	memberships["highwayMPG"] = Membership{"highwayMPG", [][]float64{{0, 0, 20, 30}, {20, 30, 40}, {30, 40, 60, 60}}}
	memberships["price"] = Membership{"price", [][]float64{{0, 0, 7000, 10000}, {7000, 10000, 20000}, {10000, 20000, 40000, 40000}}}
	return memberships
}

func testEq(a, b []float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func readFile() []Car {
	inputs := []Car{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		id, _ := conv.Atoi(strs[0])

		risk, _ := conv.ParseFloat(strs[1], 64)
		valueLoss, _ := conv.ParseFloat(strs[2], 64)
		horsepower, _ := conv.ParseFloat(strs[3], 64)
		cityMPG, _ := conv.ParseFloat(strs[4], 64)
		highwayMPG, _ := conv.ParseFloat(strs[5], 64)
		price, _ := conv.ParseFloat(strs[6], 64)
		inputs = append(inputs, Car{id, risk, valueLoss, horsepower, cityMPG, highwayMPG, price})
	}
	return inputs
}

// func loggingSetup() {
// 	logs = make(map[string]*os.File)
// }

func createFile(name string) *os.File {
	file, fileErr := os.Create(name)
	if fileErr != nil {
		print(fileErr)
		return nil
	}
	return file
}
