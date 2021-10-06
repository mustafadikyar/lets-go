package main

import (
	"fmt"
	"time"
)

func main() {
	birthDate := time.Date(2021, time.April, 27, 20, 0, 0, 0, time.UTC) //custom tarih
	fmt.Println(birthDate)

	currentDate := time.Now()
	fmt.Println(currentDate)         //o anki tarih
	fmt.Println(currentDate.Year())  //o anki yıl
	fmt.Println(currentDate.Day())   //o anki gün
	fmt.Println(currentDate.Month()) //o anki ay

	newDay := birthDate.AddDate(0, 0, 1).Day()     //+1 gün ekler
	newMonth := birthDate.AddDate(0, 1, 0).Month() //+1 ay ekler
	newYear := birthDate.AddDate(1, 0, 0).Year()   //+1 yıl ekler
	fmt.Println(newDay, newMonth, newYear)
}
