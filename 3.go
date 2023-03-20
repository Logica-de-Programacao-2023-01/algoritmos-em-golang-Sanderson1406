package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var imc float64
	fmt.Println("Insira seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira seua altura: ")
	fmt.Scan(&altura)
	altura *= altura
	imc = peso / altura
	fmt.Println("Seu IMC é: ", imc)
	if imc < 18.5 {
		fmt.Print("Você está abaixo do peso normal")
	} else if 18.5 <= imc && imc <= 24.9 {
		fmt.Print("Você está no peso ideal")
	} else if 25 <= imc && imc <= 29.9 {
		fmt.Print("Você está com sobrepeso")
	} else if imc >= 30 {
		fmt.Print("Você está com obesidade")
	}
}
