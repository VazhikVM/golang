/*Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.*/

package main

import "fmt"

func main() {
	fmt.Println(fibonacci(150))
}

func fibonacci(n int) int {
	phi := []int{1, 1, 2}

	if n > 2 {
		for i := 0; i < n; i++ {
			phi = append(phi, phi[len(phi)-1]+phi[len(phi)-2])
		}
		return phi[n-1]
	} else if n < 1 {
		println("Число не может быть меньше 1!")
		return 0
	} else {
		return 1
	}
}
