package main

import "fmt"

func main() {
	var p1, p2, p3, soma2 float64
	var peso1 float64 = 2
	var peso2 float64 = 3
	var peso3 float64 = 5
	var soma float64 = peso3 + peso2 + peso1
	fmt.Println("Número 1: ")
	fmt.Scanln(&p1)
	fmt.Println("Número 2: ")
	fmt.Scanln(&p2)
	fmt.Println("Número 3: ")
	fmt.Scanln(&p3)
	ele1 := peso1 * p1
	ele2 := peso2 * p2
	ele3 := peso3 * p3
	soma2 = ele3 + ele2 + ele1
	resul := soma2 / soma
	fmt.Println("Sua média ponderada é:", resul)
}
