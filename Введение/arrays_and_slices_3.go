package main

import "fmt"

func main() {
	fmt.Println("")
	//var N = 6
	var myArray []string

	str := "4 5 3 4 2 3"

	for _, val := range str {
		if !(string(val) == " ") {
			myArray = append(myArray, string(val))
		}
	}

	for idx, val := range str {
		if idx%2 == 0 {
			fmt.Printf("%d ", idx)
			fmt.Printf("%s ", string(val))
		}
	}

	fmt.Println(myArray)

}
