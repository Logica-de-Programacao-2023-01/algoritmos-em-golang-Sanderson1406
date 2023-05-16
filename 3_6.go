package main

import "fmt"

func main() {
	var n int
	var ta int = 10
	var ele int = 1
	fmt.Println("Insira um número ")
	fmt.Scan(&n)
	for i := 0; i <= ta; i++ {
		fmt.Println(n * ele)
		ele++
	}
}
