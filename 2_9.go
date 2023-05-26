package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int
	fmt.Println("1º Termo: ")
	fmt.Scan(&n1)
	fmt.Println("2º Termo: ")
	fmt.Scan(&n2)
	fmt.Println("3º Termo: ")
	fmt.Scan(&n3)
	resul := []int{n1, n2, n3}
	sort.Ints(resul)
	fmt.Println(resul)
}
