package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcd"
	fmt.Println("Символы уникальны? Ответ:", unique(str))

	str = "abCdefAaf"
	fmt.Println("Символы уникальны? Ответ:", unique(str))

	str = "aabcd"
	fmt.Println("Символы уникальны? Ответ:", unique(str))
}

func unique(str string) bool {
	str = strings.ToLower(str)
	strRunes := []rune(str)

	codeMap := make(map[int]struct{})

	for _, oneRune := range strRunes {
		_, ok := codeMap[int(oneRune)]
		if ok {
			return false
		} else {
			codeMap[int(oneRune)] = struct{}{}
		}
	}

	return true
}
