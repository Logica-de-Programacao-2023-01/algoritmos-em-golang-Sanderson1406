package main

import "fmt"

func main() {
	var n1 uint64
	fmt.Println("Digite um número: ")
	fmt.Scanln(&n1)
	if n1%3 == 0 && n1%5 == 0 {
		fmt.Println("O número é multiplo de 3 e 5 ao mesmo tempo")
	} else {
		fmt.Println("Nãp é multiplo de 3 e 5 ao mesmo tempo")
	}
}
