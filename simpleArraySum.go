package main

import (
	"fmt"
	"os"
)

func promptForArrayLength() int {
	fmt.Print("How long is the array? (Between 1 to 1000) ")
	reader := initReader()
	arrayLength := readLineAsInt(reader)
	if arrayLength < 0 || arrayLength >= 1000 {
		fmt.Print("Follow the rules!")
		os.Exit(1)
	}
	return arrayLength
}

// How do I enforce entered ints match the length specified in promptForArrayLength?
func populateArray(length int) []int {
	reader := initReader()
	array := readLineAsInts(reader, " ")
	if len(array) != length {
		fmt.Print("You didn't enter the correct number of array elements.")
		os.Exit(2)
	}
	return array
}

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(args []int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func promptAndSumArray() {
	fmt.Println("Let's sum an integer array!")
	length := promptForArrayLength()
	array := populateArray(length)
	sum := simpleArraySum(array)
	fmt.Print("Your sum is: ", sum)
}
