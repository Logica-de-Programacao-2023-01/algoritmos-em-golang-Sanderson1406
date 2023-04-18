package main

import "fmt"

func main() {
	n := 3
	for i := 0; i < 30; i++ {
		fmt.Println(n * i)
		i++
	}
}
