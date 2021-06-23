package main

import (
	"fmt"

	"rsc.io/quote"

	greetings "github.com/kaluva-venky/greetings-go"
)

func main() {

	// Get a greeting message and print it.
	message := greetings.Hello("Venky")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
