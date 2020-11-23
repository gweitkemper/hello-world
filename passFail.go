//This program compares an entered grade to a graing scale and displays the result
package main
import (
  "fmt"
  "log"
  "strings"
  "strconv"
  "os"
  "bufio"
)

//declaring global variables
var input string
var grade float64

func pF(){
  //order of functions
  gradeInput()
  rubric(grade)
  rerun()
}

func gradeInput (){
  //capture user input of grade
  fmt.Print("What grade did you receive?  ")
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("ERROR:")
      log.Fatal(err)
    }
  input = strings.TrimSpace(input)
  //identify captured input and convert to float for comparison
  grade, err = strconv.ParseFloat(input, 64)
    if err != nil {
      fmt.Println("ERROR:")
      log.Fatal(err)
    }
  }

func rubric(gr float64){
  //compare grade entered to grading scale
  if gr >= 60{
    fmt.Println("Congratulations! You passed.")
  } else {
    fmt.Printf("Better luck next time. You were only %.2f percentage points away from passing.\n",60-grade)
  }
}

func rerun(){
  //allow user to enter additional grades
  fmt.Println("Would you like to enter another grade? (y/n)")
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("ERROR:")
      log.Fatal(err)
    }
    input = strings.TrimSpace(input)
    if input == "y"{
      pF()
    } else {
      fmt.Println("Thanks for using passFail.go!")
    }
}
