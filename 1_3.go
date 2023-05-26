package main

import "fmt"

func main() {
	var peso, altura, imc float64
	fmt.Println("Insira seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira seua altura: ")
	fmt.Scan(&altura)
	imc = peso / (altura * altura)
	fmt.Println("Seu IMC é: ", imc)
	if imc < 18 {
		fmt.Print("Você está abaixo do peso normal")
	} else if imc >= 18 && imc <= 24 {
		fmt.Print("Você está no peso ideal")
	} else if imc >= 25 && imc <= 29 {
		fmt.Print("Você está com sobrepeso")
	} else if imc >= 30 {
		fmt.Print("Você está com obesidade")
	}
}
