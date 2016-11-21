// Oct 20, 2016

/*

Players
name                rating   take   don't take   max         cost
----------------------------------------------------- 0 ---------------0
action                8       8+0       0             8       3
bisu                  2       2+0       8             8       1
classic               4       4+8       8            12       2
dear                 10       10+8      12           18       4
effort                3       3+12      18           18       1

max(take, not take)

8 8 2 2
4 10 10 3
3 x x x

*Team = Subset of Players
TR = Sum of ratings of players on the team
BTR = Best possible TR given Players and Team rules

RULE 1: Team = subset of P, but no consecutive players
RULE 2: Total budget of 7 cost points (disregard Rule 1)
*/

// To execute Go code, please declare a func main() in a package "main"
package main

import "fmt"

var print = fmt.Println

var arr = []int{8, 2, 4, 10, 3}
var costArr = []int{3, 1, 2, 4, 1}

func main() {
	fmt.Println(maximizeBTR(0, 7))
}

func maximizeBTR(index, cost int) int {

	var K [10][10]int

	// prevDontTake := 0
	// prevMax := 0
	for i, _ := range arr[index:] {
		for c := 0; c <= cost; c++ {
			if i == 0 || c == 0 {
				K[i][c] = 0
			} else if costArr[i-1] <= c {
				K[i][c] = maxBetween(arr[i-1]+K[i-1][c-costArr[i-1]], K[i-1][c])
			} else {
				K[i][c] = K[i-1][c]
			}

			// get optimal cost from previous players
			// what is the optimal cost from not taking current player

			// dontTake := prevMax
			// take := arr[i] + prevDontTake
			// prevDontTake = prevMax
			// prevMax = maxBetween(take, dontTake)
		}

	}
	print(K)
	return K[len(arr)-1][cost]
}

func maxBetween(one, two int) int {
	if one >= two {
		return one
	} else {
		return two
	}
}

/*
   Your previous Ruby content is preserved below:


   def say_hello
     puts 'Hello, World'
     end

     5.times { say_hello }

*/
