package main

import "fmt"

func main() {
	var dias float64
	var diaria float64
	fmt.Println("Quantos dias você trabalhou? ")
	fmt.Scan(&dias)
	fmt.Println("Qual o valor da sua diária? ")
	fmt.Scan(&diaria)
	fmt.Println("Seu salário é: ", float64(dias*diaria))
}
