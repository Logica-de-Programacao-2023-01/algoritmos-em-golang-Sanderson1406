package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5, n6 int
	fmt.Println("Digite 1 termo: ")
	fmt.Scan(&n1)
	fmt.Println("Digite mais 1 termo: ")
	fmt.Scan(&n2)
	fmt.Println("Digite mais 1 termo: ")
	fmt.Scan(&n3)
	fmt.Println("Digite mais 1 termo: ")
	fmt.Scan(&n4)
	fmt.Println("Digite mais 1 termo: ")
	fmt.Scan(&n5)
	fmt.Println("Digite o ultimo termo: ")
	fmt.Scan(&n6)
	soma := n1 + n2 + n3 + n4 + n5 + n6
	resul := soma / 6
	for i := 0; i < 7; i++ {
		if n1 == 0 {
			fmt.Println("Exite um 0")
			break
		} else if n2 == 0 {
			fmt.Println("Exite um 0")
			break
		} else if n3 == 0 {
			fmt.Println("Exite um 0")
			break
		} else if n4 == 0 {
			fmt.Println("Exite um 0")
			break
		} else if n5 == 0 {
			fmt.Println("Exite um 0")
			break
		} else if n6 == 0 {
			fmt.Println("Exite um 0")
			break
		} else {
			fmt.Println("A média é: ", resul)
		}
	}

}
