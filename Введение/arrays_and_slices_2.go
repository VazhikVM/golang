/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/
package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Printf("Введите число № %d: ", i+1)
		a, _ := fmt.Scan(&a)
		array[i] = a
	}
	var maxNumber int

	/*for idx, value_array := range array {
	for_chech := value_array
	if (for_chech < 0 && idx == 0){
		max_number = for_chech
	}*/
	for idx, valueArray := range array {

		if valueArray < 0 && idx == 0 {
			maxNumber = valueArray
		} else if valueArray > maxNumber {
			maxNumber = valueArray
		}
	}
	fmt.Printf("Максимальное число массива %d", maxNumber)
}
