package main

import "fmt"

func main() {
	var n uint64
	fmt.Println("Insira um número: ")
	fmt.Scan(&n)
	fmt.Println("Seu sucessor é: ", n+1)
	fmt.Println("Seu antecessor é: ", n-1)
}
