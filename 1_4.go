package main

import "fmt"

func main() {
	var p1, p2, p3 float64
	var x1 float64 = 2
	var x2 float64 = 3
	var x3 float64 = 5
	var ele1 float64 = 0
	var ele2 float64 = 0
	var ele3 float64 = 0
	var baixo float64 = 0
	fmt.Println("Número 1: ")
	fmt.Scanln(&p1)
	fmt.Println("Número 2: ")
	fmt.Scanln(&p2)
	fmt.Println("Número 3: ")
	fmt.Scanln(&p3)
	ele1 = x1 * p1
	ele2 = x2 * p2
	ele3 = x3 * p3
	baixo = p1 + p2 + p3
	fmt.Println("Sua média ponderada é:", ele1+ele2+ele3/baixo)
}
