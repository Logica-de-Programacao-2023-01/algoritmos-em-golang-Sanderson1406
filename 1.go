package main

import "fmt"

func main() {
	var n1 int32
	var n2 int32
	var n3 int32
	fmt.Print("Número 1: ")
	fmt.Scan(&n1)
	fmt.Print("Número 2: ")
	fmt.Scan(&n2)
	fmt.Print("Número 3: ")
	fmt.Scan(&n3)
	fmt.Println("A soma desses números =", n1+n2+n3)
}
