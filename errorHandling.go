package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	//
	_, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal("Hata : ", err)
		fmt.Println(err)
	}

	//
	myError := errors.New("Bu bir hata!")
	fmt.Println(myError.Error())
}
