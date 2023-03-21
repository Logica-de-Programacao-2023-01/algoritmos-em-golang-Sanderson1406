package main

import "fmt"

func main() {
	var n1 uint64
	var n2 uint64
	fmt.Println("Insira 1 número: ")
	fmt.Scanln(&n1)
	fmt.Println("Insira outro número: ")
	fmt.Scanln(&n2)
	if n1 == n2 {
		fmt.Println("O primeiro número é igual ao segundo número")
	} else if n1 > n2 {
		fmt.Println("O primeiro número é maior que o segundo número")
	} else if n1 < n2 {
		fmt.Println("O primeiro número é menor que segundo número")
	}
}
