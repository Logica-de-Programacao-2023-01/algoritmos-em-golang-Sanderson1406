package main

import "fmt"

func main() {
	i := 10
	for i <= 10 {
		fmt.Println(i)
		i--
		if i == 0 {
			break
		}
	}
}
