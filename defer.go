package main

import "fmt"

var isConnect bool = false

func main() {
	//defer, içerisindeki fonksiyondaki tüm işlemler bitene kadar ilgili işlemi bekletir.
	defer fmt.Println("Hello")

	fmt.Println("One")
	fmt.Println("Two")
	fmt.Println("Three")

	/*
		One
		Two
		Three
		Hello
	*/

	fmt.Printf("connection open : %v\n", isConnect)
	databaseProcessing()
	fmt.Printf("connection open : %v\n", isConnect)

	/*
		connection open : false
		Connect to DB
		Deferring disconnect!
		connection open : true
		Doing something
		Disconnected!
		connection open : false
	*/
}

func databaseProcessing() {
	connect()
	fmt.Println("Deferring disconnect!")
	defer disconnect()
	fmt.Printf("connection open : %v\n", isConnect)
	fmt.Println("Doing something")
}

func connect() {
	isConnect = true
	fmt.Println("Connect to DB")
}

func disconnect() {
	isConnect = false
	fmt.Println("Disconnected!")
}
