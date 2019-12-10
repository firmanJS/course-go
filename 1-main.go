// for build apps
package main

import "fmt"

func main(){
	// var card = "this is card"
	// card := "this is card"
	// card = "open"

	card := newCard()

	fmt.Println(card)
	// fmt.Println("Hi There")
}

func newCard() string{
	return "this is cards"
}

