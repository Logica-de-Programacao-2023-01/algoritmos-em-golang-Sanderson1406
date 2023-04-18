package main

import "fmt"

func main() {
	var n1, n2, n3 uint32
	fmt.Println("1º Termo: ")
	fmt.Scan(&n1)
	fmt.Println("2º Termo: ")
	fmt.Scan(&n2)
	fmt.Println("3º Termo: ")
	fmt.Scan(&n3)
	if n1 < n2 && n1 < n3 {
		fmt.Println(n1)
	} else if n2 < n1 && n2 > n3 {
		fmt.Println(n2)
	} else if n2 > n1 && n2 < n3 {
		fmt.Println(n2)
	}
	if n3 < n1 && n3 < n2 {
		fmt.Println(n2)
	}
	if n2 < n1 && n2 > n3 {
		fmt.Println(n2)
	}
	if n2 < n1 && n2 > n3 {
		fmt.Println(n2)
	}
	if n2 < n1 && n2 > n3 {
		fmt.Println(n2)
	}
	if n2 < n1 && n2 > n3 {
		fmt.Println(n2)
	}

}
