// To execute Go code, please declare a func main() in a package "main"

// 1000 => "1,000"
// 1234567 => "1,234,567"

package main

import "fmt"
import "strconv"
import "strings"

var print = fmt.Println

func main() {
	commafy(1234)
	print(commafy(1234567))
	print(commafy(12345678))
	print(commafy(123456789))
	print(commafy(0))
}

func commafy(num int) string {
	arr := strings.Split(strconv.Itoa(num), "")
	// print(arr)

	counter := 0
	result := ""
	for i := len(arr) - 1; i >= 0; i-- {
		// print(arr[i])
		if counter%3 == 0 && counter != 0 {
			result = "," + result
		}
		result = arr[i] + result

		// print(result)
		counter += 1
	}

	return result
}
