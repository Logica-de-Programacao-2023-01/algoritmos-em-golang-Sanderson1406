package main

import "fmt"

func main() {
	var kls float64
	fmt.Println("Insira seu peso em KLs: ")
	fmt.Scanln(&kls)
	libras := 2.2
	fmt.Println("Seu peso em Libras é: ", kls*libras)
}
