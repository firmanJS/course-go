package main
// import "fmt"

// # read file
// func main(){
// 	card := newDeckFromFile("my_cards")
// 	card.print()
// }

// # save file
// func main(){
// 	card := newDeck()
// 	card.saveTofile("my_cards")
// }

// # join or split string print
// func main(){
// 	card := newDeck()
// 	fmt.Println(card.toString())
// }

// # shuffle
func main(){
	card := newDeck()
	card.shuffle()
	card.print()
}