package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

//constructure
func CreateHuman() *Human {
	human := new(Human)
	return human
}

func main() {
	human := CreateHuman()
	human.Name = "Mustafa"
	fmt.Println(human.Name)

	mustafa := Human{Name: "Mustafa", Surname: "Dikyar", Age: 27}
	fmt.Println(mustafa.Age) //27

	ahmet := new(Human)
	ahmet.Age = 33

	fmt.Println(ahmet.Age) //33
}
