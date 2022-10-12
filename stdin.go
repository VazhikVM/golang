package main

import "fmt"
import "math"


func main () {
	var num float64;
	var ex string;

	for {
		fmt.Println("Введите десятичное число: ")
		fmt.Scan(&num)

		if (num <= 0){
			fmt.Printf("Число R не подходи%2.2f\n", num)
		} else if (num >10000) {
			fmt.Printf("Число R не подходит %e\n", num)
		} else {
			fmt.Printf("%.4f\n", math.Pow(num, 2))
		}

		fmt.Println("Для выхода введете 'exit' и нажмите enter: ")
		fmt.Scan(&ex)

		if (ex == "exit") {
			break
		}
	}
}