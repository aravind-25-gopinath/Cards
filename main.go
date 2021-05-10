package main

import "fmt"

func main() {
	//var card string = "Ace of Spades" //more explicit of way telling Go that 'card' is a String variable
	card := "Ace of Spades" //this lets Go infer the data type of the variable (:= only used when we FIRST instantiate a variable)
	fmt.Println(card)
}
