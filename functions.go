package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello("Mustafa")
	add(1, 4)
	name, _ := getNameAndSurname("Mustafa Dikyar")
	fmt.Println(name)

	sentence := combine("Bu", "bir", "cümledir.")
	fmt.Println(sentence)

	name, surname := split("Mustafa Dikyar")
	fmt.Println(name, surname)

	anonymous := func(name string) (fullMessage string) {
		fullMessage = "Hello " + name
		return
	}

	message := anonymous("Mustafa")
	fmt.Println(message)
}

//Usage 1
func sayHello(name string) {
	fmt.Println("Hello " + name)
}

//Usage 2
func add(number1, number2 int) int {
	return number1 + number2
}

//Usage 3
func getNameAndSurname(fullName string) (string, string) {
	splitted := strings.Split(fullName, " ")
	return splitted[0], splitted[1]
}

//Usage 4
func combine(terms ...string) string { //değişken sayıda parametre alması durumu
	var sentence string
	for _, term := range terms {
		sentence += term + " "
	}

	return sentence
}

//Usage 5
func split(fullName string) (name, surname string) {
	splitted := strings.Split(fullName, " ")
	name = splitted[0]
	surname = splitted[1]
	return
}
