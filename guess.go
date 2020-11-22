//Guess is a game in which the user attempts to guess a random number generated by the program
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guessingGame() {
	//Introduce the game
	fmt.Println("I'm thinking of a number between 1 and 100, can you guess it?")
	//Generate a random number between 1 and 100
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I'll give you 10 tries.")
  //Capture input
  reader := bufio.NewReader(os.Stdin)
	//Allow user to guess up to 10 times
  for guesses := 10; guesses == 0; guesses-- {
    fmt.Println("What's your guess?")
    input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR:")
			log.Fatal(err)
		}
		//Convert input to integer
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
      if err != nil{
        fmt.Println("ERROR:")
        log.Fatal(err)
      }
		//Compare guess to target number
		if guess == target {
			fmt.Println("What're you? Psychic? Nice work. I quit.")
			
		} else if guesses == 0{
      fmt.Println("Looks like you're out of guesses. Better luck next time. The number was", target,".")
      break
    } else if guess > target {
			fmt.Println("Too high! You have", guesses,"left.")
		} else if guess < target {
			fmt.Println("Too low! You have", guesses,"left.")
		} else{
			fmt.Println("Hmm... something went wrong. Guess again.")
		}
	}
}