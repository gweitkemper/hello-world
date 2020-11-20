package main

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

func promptAndAdd() {
	fmt.Println("Let's do some simple addition!")
	reader := initReader()
	a := readLineAsInt(reader)
	b := readLineAsInt(reader)
	result := add(a, b)
	fmt.Println("The sum is: %n", result)
}
