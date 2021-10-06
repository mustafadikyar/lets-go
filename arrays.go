package main

import "fmt"

func main() {

	//Usage 1
	cities := [3]string{}
	cities[0] = "Antalya"
	cities[1] = "İstanbul"
	cities[2] = "İzmir"

	fmt.Println(cities) //get

	cities[2] = "Ankara" //set
	fmt.Println(cities[2])

	//Usage 2
	numbers := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(numbers)

	//Usage 3
	var colors [3]string
	colors[0] = "Sarı"
	colors[1] = "Mavi"
	colors[2] = "Yeşil"

	i := 0
	for i <= len(colors)-1 {
		fmt.Println(colors[i])
		i++
	}
}
