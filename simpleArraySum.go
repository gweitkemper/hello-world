package main

import "fmt"

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(args []int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum = sum + args[i]
	}
	return sum
}

func promptAndSumArray() {
	fmt.Println("Let's sum an integer array!")
	reader := initReader()
	result := simpleArraySum(readLineAsInts(reader, " "))
	// var stdOut = initFile("/dev/repo/input/simpleArray.json")
	// writeFile(stdOut, string(result))
	fmt.Println("The sum is:", result)
}
