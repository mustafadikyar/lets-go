package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	for i := range numbers {
		fmt.Println("index", i, "value", numbers[i])
	}

	/*
		index 0 value 10
		index 1 value 20
		index 2 value 30
		index 3 value 40
		index 4 value 50
	*/

	capitals := map[string]string{"Türkiye": "Ankara", "İspanya": "Madrid", "Fransa": "Paris", "İtalya": "Roma"}
	for key, val := range capitals {
		fmt.Println("key", key, "value", val)
	}

	/*
		key Türkiye value Ankara
		key İspanya value Madrid
		key Fransa value Paris
		key İtalya value Roma
	*/
}
