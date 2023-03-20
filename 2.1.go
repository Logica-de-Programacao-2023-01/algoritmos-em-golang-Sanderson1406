package main

import "fmt"

func main() {
	var n1 uint64
	var n2 uint64
	fmt.Println("Insira 1 número: ")
	fmt.Scanln(&n1)
	fmt.Println("Insira outro número: ")
	fmt.Scanln(&n2)
	fmt.Println("O primeiro número == segundo número é ", n1 == n2)
	fmt.Println("O primeiro número > segundo número é ", n1 > n2)
	fmt.Println("O primeiro número < segundo número é ", n1 < n2)
	fmt.Println("O primeiro número != segundo número é ", n1 != n2)
}
