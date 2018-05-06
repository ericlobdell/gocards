package main

func main() {
	cards := deck{"Ace of diamonds", "Ace of spades"}
	cards = append(cards, "King of hearts")

	cards.print()
}
