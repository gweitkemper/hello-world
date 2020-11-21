package main

import "fmt"

func multiply(a int, b int) int {
	return (a * b)
}

func promptAndMultiply() {
	fmt.Println("Let's do some simple multiplication!")
	reader := initReader()
	a := readLineAsInt(reader)
	b := readLineAsInt(reader)
	result := multiply(a, b)
	fmt.Println("The product is:", result)
}