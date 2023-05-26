package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número: ")
	fmt.Scan(&n)
	for i := 1; i > 0; i++ {
		if n%i == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}
