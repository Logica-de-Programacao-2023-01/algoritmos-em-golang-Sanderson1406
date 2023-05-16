package main

import "fmt"

func main() {
	var id uint32
	fmt.Println("Qual a sua idade?")
	fmt.Scan(&id)
	if id <= 9 {
		fmt.Println("Você é um Mirim")
	} else if id >= 10 && id <= 13 {
		fmt.Println("Você é Infantil")
	} else if id >= 14 && id <= 17 {
		fmt.Println("Você é um Juvenil")
	} else if id >= 18 {
		fmt.Println("Você é um Adulto")
	}
}
