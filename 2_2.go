package main

import "fmt"

func main() {
	var n1, n2, n3 uint32
	fmt.Println("Insira um número: ")
	fmt.Scanln(&n1)
	fmt.Println("Insira outro número: ")
	fmt.Scanln(&n2)
	fmt.Println("Insira mais um número: ")
	fmt.Scanln(&n3)
	if n1 < n2 && n1 < n3 {
		fmt.Println(n1, "é o menor")
	} else if n2 < n1 && n2 < n3 {
		fmt.Println(n2, "é o menor")
	} else if n3 < n1 && n3 < n2 {
		fmt.Println(n3, "é o menor")
	}
}
