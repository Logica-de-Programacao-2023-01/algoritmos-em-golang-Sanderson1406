package main

import "fmt"

func main() {
	var n1 int64
	var n2 int64
	fmt.Println("1º Termo: ")
	fmt.Scan(&n1)
	fmt.Println("2º Termo: ")
	fmt.Scan(&n2)
	if n1 < 1 && n2 < 1 {
		fmt.Println("O resultado da multiplicação desses termos é: ", n1*n2)
	} else if n1 < -1 || n2 < -1 {
		fmt.Println("O resultado da soma desses termos é: ", n1-n2)
	}
}
