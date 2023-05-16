package main

import "fmt"

func main() {
	var valor float64
	var valor2 float64
	var valor3 float64
	fmt.Println("Qual o valor do produto? ")
	fmt.Scan(&valor)
	valor2 = valor * 10 / 100
	valor3 = valor - valor2
	fmt.Println("Seu valor com desconto é: ", valor3)
}
