package main

func main() {
	//var card string = "Ace of Spades" //more explicit of way telling Go that 'card' is a String variable
	/*card := newCard() //this lets Go infer the data type of the variable (:= only used when we FIRST instantiate a variable)
	fmt.Println(card)*/

	//cards := []string{newCard(), newCard()} //a slice of 2 string elements
	cards := deck{"Ace of Diamonds", newCard()} //deck is something that is initialized in deck.go, cards is of type deck
	cards = append(cards, "Six of Spades")      //adds "Six of Spades" to the slice

	//iterate over the cards slice
	/*for i, card := range cards {
		fmt.Println(i, card)
	}

	//another way of iterating over the cards
	for i := 0; i < len(cards); i++ {
		fmt.Println(i, cards[i])
	}*/

	cards.print() //reciever function from deck.go
}

//function that returns type string ("Five of Diamonds")
func newCard() string {
	return "Five of Diamonds"
}
