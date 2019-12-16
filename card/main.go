package main

func main(){
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// card = "Five of Diamonds"
	card := deck{"Ace of Spades", newCard()} 
	card = append(card, "Six of Spades")
	card.print()
}

func newCard() string {
	return "Five of Diamonds"
}