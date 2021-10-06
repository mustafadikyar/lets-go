package main

import "fmt"

func main() {
	language := "de"

	switch language {
	case "tr":
		fmt.Println("Türkçe")
	case "en":
		fmt.Println("İngilizce")
	case "ru":
		fmt.Println("Rusça")
	default:
		fmt.Println("Diğer")
	}
}
