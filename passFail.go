//This program compares an entered grade to a graing scale and displays the result
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passFail() {
	//order of functions
	grade := gradeInput()
	rubric(grade)
	rerun()
}

func gradeInput() float64 {
	//capture user input of grade
	fmt.Print("What grade did you receive?  ")
	input := getInput()
	//identify captured input and convert to float for comparison
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	return output
}

func rubric(grade float64) {
	var gradeA float64 = 94
	var gradeB float64 = 86
	var gradeC float64 = 74
	var gradeD float64 = 65

	//compare grade entered to grading scale
	switch {
	case grade >= gradeA:
		fmt.Println("Congratulations! You got an A.")
	case grade >= gradeB:
		fmt.Printf("Not bad. You got a B and were only %.2f percentage points away from an A.\n", gradeA-grade)
	case grade >= gradeC:
		fmt.Printf("Meh. You got a C and were only %.2f percentage points away from a B.\n", gradeB-grade)
	case grade >= gradeD:
		fmt.Printf("Better luck next time. You were only %.2f percentage points away from a C.\n", gradeC-grade)
	default:
		fmt.Printf("Congratulations! Despite being %.2f percentage points away from passing, you've received the greatest gift of all: Failure.\n", gradeD-grade)
	}
}

func rerun() {
	//allow user to enter additional grades
	fmt.Println("Would you like to enter another grade? (y/n)")
	input := getInput()
	if input == "y" {
		passFail()
	} else {
		fmt.Println("Thanks for using passFail.go!")
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}
