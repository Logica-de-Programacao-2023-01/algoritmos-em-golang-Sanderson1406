package main

import "fmt"

func main() {
	var n1 uint32
	var n2 uint32
	var n3 uint32
	fmt.Println("Insira um número: ")
	fmt.Scanln(&n1)
	fmt.Println("Insira outro número: ")
	fmt.Scanln(&n2)
	fmt.Println("Insira mais um número: ")
	fmt.Scanln(&n3)
	fmt.Println("O primeiro é menor que o segundo? ", n1 < n2)
	fmt.Println("O primeiro é menor que o terceiro? ", n1 < n3)
	fmt.Println("O segundo é menor que o terceiro? ", n2 < n3)
	fmt.Println("O segundo é menor que o primeiro? ", n2 < n1)
	fmt.Println("O terceiro é menor que o primeiro? ", n3 < n1)
	fmt.Println("O terceiro é menor que o segundo? ", n3 < n2)

}
