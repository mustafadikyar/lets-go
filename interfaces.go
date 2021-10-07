package main

import "fmt"

type product struct {
	name string
}

type iProductRepository interface {
	getProductName() string
}

func (product product) getProductName() string {
	return product.name
}

func printProductName(product iProductRepository) {
	fmt.Println(product.getProductName())
}

func main() {
	product := product{name: "ürün 1"}
	printProductName(product)
}
