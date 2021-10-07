package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	mySlice := numbers[:] //: tüm elemanların sayısını ifade eder.
	fmt.Println(mySlice)  //[1, 2, 3, 4, 5, 6, 7, 8, 9, 0]

	/*Slice veri tutmaz. Slice üzerinde yapılan bir değişiklik diziyi etkiler.*/
	mySlice[2] = 11
	fmt.Println(numbers) //[1, 2, 11, 4, 5, 6, 7, 8, 9, 0]

	numbers = append(numbers, 10) //veri ekler
	fmt.Println(numbers)          //[1 2 11 4 5 6 7 8 9 0 10]

	//dizinin bir kesitini alma.
	mySlice2 := numbers[2:6]
	fmt.Println(mySlice2) //[11 4 5 6]
}
