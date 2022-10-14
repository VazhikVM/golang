package main

import "fmt"

func main() {
	fmt.Println(maxCountNum(1, 0, 0))
}
func maxCountNum(x, y, z int) int {
	if x+y+z > 1 {
		return 1
	}
	return 0
}
