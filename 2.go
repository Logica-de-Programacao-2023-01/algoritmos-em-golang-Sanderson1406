package main

import "fmt"

func main() {
	var n uint32
	fmt.Println("Insira um número:")
	fmt.Scan(&n)
	fmt.Println("Seu dobro é: ", n*2)
	fmt.Println("Seu triplo é: ", n*3)
	fmt.Println("Seu quadruplo é: ", n*4)
}
