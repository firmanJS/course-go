package main

func main() {
	// # receiver function
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// card = "Five of Diamonds"
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()

	// cards := newDeck()
	// cards.print()

	// # multiple return
	cards := newDeck()
	hand, RemainingCards := deal(cards,5)
	hand.print()
	RemainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
