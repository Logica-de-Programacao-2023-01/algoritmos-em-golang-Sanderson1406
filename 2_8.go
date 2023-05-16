package main

import "fmt"

func main() {
	var dia int32
	fmt.Println("Qual a contagem do dia em números? ")
	fmt.Scan(&dia)
	if dia == 1 {
		fmt.Println("O dia é domingo")
	} else if dia == 2 {
		fmt.Println("O dia é segunda-feira")
	} else if dia == 3 {
		fmt.Println("O dia é terça-feira")
	} else if dia == 4 {
		fmt.Println("O dia é quarta-feira")
	} else if dia == 5 {
		fmt.Println("O dia é quinta-feira")
	} else if dia == 6 {
		fmt.Println("O dia é sexta-feira")
	} else if dia == 7 {
		fmt.Println("O dia é sabádo")
	}
}
