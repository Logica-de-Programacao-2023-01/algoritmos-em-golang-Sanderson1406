package main

import "fmt"

func main() {
	var sala float64
	var mais float64 = 15 / 100
	fmt.Println("Qual é o seu sálario? ")
	fmt.Scan(&sala)
	fmt.Println("Seu novo sálario é: ", sala+mais)
}
