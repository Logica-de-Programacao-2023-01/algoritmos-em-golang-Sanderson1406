package main

import "fmt"

func main() {
	var sala uint64
	var valor1 uint64
	var nsala uint64
	fmt.Println("Qual é o seu sálario? ")
	fmt.Scan(&sala)
	valor1 = sala * 15 / 100
	nsala = sala + valor1
	fmt.Println("Seu novo sálario é: ", nsala)
}
