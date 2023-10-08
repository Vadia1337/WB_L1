package main

import (
	"fmt"
)

func main() {
	str := "snow dog sun"
	runeStr := []rune(str)

	reversedStr := make([]rune, len(runeStr))

	wordBuff := make([]rune, 0)

	for i := 0; i < len(runeStr); i++ {
		wordBuff = append(wordBuff, runeStr[i])

		// проверка на пробел для перестановки в начало слова
		if runeStr[i] == 32 {
			wordBuff = append(wordBuff[len(wordBuff)-1:], wordBuff[:len(wordBuff)-1]...)
		}

		// проверяем на пробел или на конец цикла и добавляем в выходную строку
		if runeStr[i] == 32 || i == len(runeStr)-1 { // пробел в разных кодировках (utf-8, unicode) имеет шифр 32
			reversedStr = append(wordBuff, reversedStr...)
			wordBuff = nil
		}

	}

	fmt.Println(string(reversedStr))

}
