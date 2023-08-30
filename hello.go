package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var words = words()
	var rNum = rNum()
	fmt.Println(words[rNum])
}

func words() [20]string {
	var words [20]string

	words[0] = "Yes, definitely"
	words[1] = "No, not this time"
	words[2] = "Outlook is hazy, try again"
	words[3] = "It's a possibility"
	words[4] = "Absolutely yes"
	words[5] = "Can't say for sure"
	words[6] = "Ask again later"
	words[7] = "No doubt about it"
	words[8] = "Most likely"
	words[9] = "Better not tell you now"
	words[10] = "Not looking good"
	words[11] = "Yes, but with caution"
	words[12] = "Signs point to yes"
	words[13] = "Unlikely"
	words[14] = "You can count on it"
	words[15] = "Reply hazy, try again"
	words[16] = "Definitely no"
	words[17] = "Chances are slim"
	words[18] = "Ask someone else"
	words[19] = "It's up to you"

	return words
}

func rNum() int {
	randomNumber := rand.Intn(3)
	return randomNumber
}
