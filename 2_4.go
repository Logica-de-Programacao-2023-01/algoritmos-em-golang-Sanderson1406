package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var sexo string
	var imc float64 = 0
	fmt.Println("Insira seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira sua altura: ")
	fmt.Scan(&altura)
	fmt.Println("Insira seu sexo: ")
	fmt.Scan(&sexo)
	imc = float64(uint64(peso / (altura * altura)))
	fmt.Println("Seu IMC é: ", imc)
	if imc == 18 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 19 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 20 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 21 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 22 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 23 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 24 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 25 && sexo == "Masculino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc < 18 && sexo == "Masculino" {
		fmt.Println("Está fora do recomendado, CUIDADO!!! ")
	} else if imc > 25 && sexo == "Masculino" {
		fmt.Println("Está acima do recomendado, CUIDADO!!! ")
	} else if imc == 18 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 19 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 20 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 21 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 22 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 23 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc == 24 && sexo == "Feminino" {
		fmt.Println("Está dentro do recomendado ")
	} else if imc < 18 && sexo == "Feminino" {
		fmt.Println("Está fora do recomendado, CUIDADO!!! ")
	} else if imc > 24 && sexo == "Feminino" {
		fmt.Println("Está acima do recomendado, CUIDADO!!! ")
	}
}
