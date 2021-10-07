package main

import "fmt"

func main() {
	//KeyValuePair
	cities := make(map[int]string)
	fmt.Println(cities) //map[]

	cities[01] = "Adana" //set
	cities[07] = "Antalya"
	cities[34] = "İstanbul"
	fmt.Println(cities)     //map[1:Adana 7:Antalya 34:İstanbul]
	fmt.Println(cities[34]) //İstanbul

	delete(cities, 34)
	fmt.Println(cities) //map[1:Adana 7:Antalya]

	for key, value := range cities {
		fmt.Printf("%v: %v\n", key, value)
	}
	/*
		1: Adana
		7: Antalya
	*/
}
