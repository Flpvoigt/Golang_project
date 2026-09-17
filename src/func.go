package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func (p Product) DisplayInfo() {
	fmt.Println("Product Name:", p.Name)
	fmt.Println("Product Price:", fmt.Sprintf("R$ %.2f", p.Price))
}
