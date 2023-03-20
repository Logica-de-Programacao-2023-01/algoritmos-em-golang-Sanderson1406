package main

import "fmt"

func main() {
	var anos float64
	var dias float64 = 0
	fmt.Println("Insira o ano: ")
	fmt.Scan(&anos)
	dias = anos * 365
	fmt.Println("Em dias é: ", dias)
}
