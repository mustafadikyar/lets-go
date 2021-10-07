package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "st"))         //true
	fmt.Println(strings.HasPrefix("go_lang", "go"))     //true
	fmt.Println(strings.HasSuffix("golang.rar", "rar")) //true
	fmt.Println(strings.Index("golang", "a"))           //3
}
