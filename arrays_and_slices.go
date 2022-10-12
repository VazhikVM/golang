package main

import "fmt"


/*func main() {
	var workArray[10]int = [10]int{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}
	fmt.Println(workArray)

	var num_1, num_2, num_3, num_4, num_5, num_6 int = 1, 2, 3, 5, 7, 8
	fmt.Println(num_1, num_2, num_3, num_4, num_5, num_6)
	
	fmt.Printf("%d %d %d %d %d %d",workArray[num_1], workArray[num_2], workArray[num_3], workArray[num_4], workArray[num_5], workArray[num_6])

}*/

func main() {
	var N, N_array int;
	fmt.Println("Введите число от 1 до 3: ")
	fmt.Scan(&N);
	new_slice := []int{}

	for i := 0; i < N; i++{
		fmt.Printf("Введите элемент %d: ", i+1)
		fmt.Scan(&N_array);

		new_slice = append(new_slice, N_array)
	}

	fmt.Println(new_slice)

}