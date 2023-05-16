package main

import "fmt"

func main() {
	var sala float64
	var valor1 float64 = 0
	var nsala float64 = 0
	fmt.Println("O seu salário atual é: ")
	fmt.Scan(&sala)
	if sala < 1000 {
		valor1 = sala * 10 / 100
		nsala = sala + valor1
		fmt.Println("Seu novo sálario é: ", nsala)
	} else if sala >= 1000 {
		valor1 = sala * 5 / 100
		nsala = sala + valor1
		fmt.Println("Seu novo sálario é: ", nsala)
	}
}
