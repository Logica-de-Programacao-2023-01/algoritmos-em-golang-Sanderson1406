package main

import "fmt"

func main() {
	var n1, n2 int64
	fmt.Println("1º Termo: ")
	fmt.Scan(&n1)
	fmt.Println("2º Termo: ")
	fmt.Scan(&n2)
	if n1 >= 0 && n2 >= 0 {
		fmt.Println("O resultado da multiplicação desses termos é: ", n1*n2)
	} else if n1 <= 0 {
		fmt.Println("O resultado da soma desses termos é: ", n1+n2)
	} else if n2 <= 0 {
		fmt.Println("O resultado da soma desses termos é: ", n1+n2)
	}
}
