package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Println("Insira seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira seua altura: ")
	fmt.Scan(&altura)
	altura *= altura
	fmt.Println("Seu IMC é: ", peso/altura)
}
