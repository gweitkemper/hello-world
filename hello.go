package main

import (
	"fmt"

	"rsc.io/quote"
)

func hello() {
	fmt.Println("Let's generate a quote!")
	fmt.Println(quote.Go())
}
