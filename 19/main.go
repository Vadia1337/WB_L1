package main

import "log"

func main() {

	strToinvert := "абырвалг"

	runeStr := []rune(strToinvert)
	invertedRune := make([]rune, len(runeStr))
	for i, v := range runeStr {
		invertedRune[len(runeStr)-1-i] = v
	}

	log.Print(string(invertedRune))
}
