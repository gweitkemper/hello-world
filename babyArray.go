package main

import (
	"fmt"
	"os"
)

func promptForArrayLength() int {
	var ar int
	fmt.Print("How long is the array? (Must be between 1 to 1000) ")
	_, err := fmt.Scanf("%d", &ar)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else if ar < 0 && ar <= 1000 {
		fmt.Print("Follow the rules!")
		os.Exit(2)
	}
	return ar
}

func populateArray(arraySize int) []int {
	fmt.Printf("Please enter %d integers separated by spaces: ", arraySize)
	array := make([]int, arraySize)
	for i := 0; i < len(array); i++ {
		fmt.Scanf("%d", &array[i])
	}
	return array
}

func babyArray() {
	// Would this be batter as populateArray(promptForArrayLength())?
	arLen := promptForArrayLength()
	array := populateArray(arLen)
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Print(sum)
}
