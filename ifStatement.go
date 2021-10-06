package main

import "fmt"

func main() {
	number := 10

	if number > 0 {
		fmt.Println("Sayı pozitif")
	} else if number < 0 {
		fmt.Println("Sayı negatif")
	} else {
		fmt.Println("Sayı sıfıra eşit")
	}
}
