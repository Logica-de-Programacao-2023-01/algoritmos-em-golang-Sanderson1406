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
	if n1 > n2 && n1 > n3 {
		fmt.Println(n1, "É maior")
	}
}
