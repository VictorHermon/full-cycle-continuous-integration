package main

import "fmt"

func main() {
	fmt.Println("O resultado da soma é:", Soma(5, 10))
}

func Soma(a int, b int) int {
	return a + b
}
