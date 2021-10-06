package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber = "10"

	//Convert

	//string to int
	number, _ := strconv.Atoi(firstNumber)
	result := number + 3
	fmt.Println(result) //13

	//int to string
	fmt.Println("Sonuç : " + strconv.Itoa(result)) //Sonuç : 13

	//Casting
	var secondNumber float64 = 11.1
	var intNumber = int(secondNumber)
	fmt.Println(intNumber) //11
}
