package main

import "fmt"

type Language string
const(
	Turkish Language = "tr"
	English Language = "en"
	Russian Language = "ru"
)

func GetLanguage(language Language){
	fmt.Println(language)
}

func main(){
	GetLanguage(Turkish) //tr
	GetLanguage(Russian) //ru
}