package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Println("Insira seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira seua altura: ")
	fmt.Scan(&altura)
	imc := altura / peso
	fmt.Println("Seu IMC é: ", imc)
	if imc < 18 {
		fmt.Print("Você está abaixo do peso normal")
	} else if 18 <= imc && imc <= 24 {
		fmt.Print("Você está no peso ideal")
	} else if 25 <= imc && imc <= 29 {
		fmt.Print("Você está com sobrepeso")
	} else if imc >= 30 {
		fmt.Print("Você está com obesidade")
	}
}
