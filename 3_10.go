package main

import "fmt"

func main() {
	var num, maior int
	fmt.Println("Digite uma sequência de números:")
	for {
		fmt.Print("Digite um número" +
			": ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > maior {
			maior = num
		}
	}
	if maior != 0 {
		fmt.Println("O maior número é:", maior)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
