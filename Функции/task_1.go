package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour(4, 5, 6, 7))
}
func minimumFromFour(a, b, c, d int) int {
	newArray := [...]int{a, b, c, d}
	result := newArray[0]
	for _, val := range newArray {
		if result > val {
			result = val
		}
	}
	return result
}
