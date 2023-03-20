package main

import "fmt"

func main() {
	var dias int64
	var diaria int64
	fmt.Println("Quantos dias você trabalhou? ")
	fmt.Scan(&dias)
	fmt.Println("Qual o valor da sua diária? ")
	fmt.Scan(&diaria)
	fmt.Println("Seu salário é: ", dias*diaria)
}
