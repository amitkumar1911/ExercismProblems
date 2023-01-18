package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.

var cards = map[string]int{"ace":11, "two":2, "three":3, "four":4, "five":5, "six":6, "seven":7, "eight":8, "nine":9, "ten":10, "jack":10,"queen":10,"king":10, "other":0}

func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")

    return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function") 

    value:=cards[card1]+cards[card2]  

    if card1=="ace" && card2=="ace" {return "P"}
    if value==21 && (dealerCard!="ace"  &&  dealerCard!="jack" && dealerCard!="queen" && dealerCard!="king" && dealerCard!="ten") {return "W"}
    if value==21 && (dealerCard=="ace" || dealerCard=="jack" || dealerCard=="queen" || dealerCard=="king"|| dealerCard=="ten") {return "S"}
    if value>=17 && value<=20 {return "S"}
    if (value>=12 && value<=16) && cards[dealerCard]<7  {return "S"} 
    if (value>=12 && value<=16) && cards[dealerCard]>=7  {return "H"} 
    return "H"
        
}
