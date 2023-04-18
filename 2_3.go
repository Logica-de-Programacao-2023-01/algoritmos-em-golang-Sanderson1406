package main

import "fmt"

func main() {
	var n uint32
	fmt.Println("Insira um número: ")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n, " é par")
	} else {
		fmt.Println(n, " é impar")
	}
}
