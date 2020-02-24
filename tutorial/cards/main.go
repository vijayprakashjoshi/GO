package main

func main() {

	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Ace of hearts")

	cards.printDeck()

	//	Printing using the print method in file deck
	//	for i, card := range cards {
	//		fmt.Println(i, card)
	//	}

}

func newCard() string {

	return "Ace of Spades"

}
