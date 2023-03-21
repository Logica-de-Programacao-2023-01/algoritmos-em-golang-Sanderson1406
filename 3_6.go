package main

import "fmt"

func main() {
	var n uint32
	fmt.Println("Insira um número ")
	fmt.Scan(&n)
	for n != 0 {
		fmt.Println(n)
		fmt.Println(n * 2)
		fmt.Println(n * 3)
		fmt.Println(n * 4)
		fmt.Println(n * 5)
		fmt.Println(n * 6)
		fmt.Println(n * 7)
		fmt.Println(n * 8)
		fmt.Println(n * 9)
		fmt.Println(n * 10)
		break
	}
}
