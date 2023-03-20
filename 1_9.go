package main

import "fmt"

func main() {
	var valor uint64
	var valor2 uint64
	var valor3 uint64
	fmt.Println("Qual o valor do produto? ")
	fmt.Scan(&valor)
	valor2 = valor * 10 / 100
	valor3 = valor - valor2
	fmt.Println("Seu valor com desconto é: ", valor3)
}
