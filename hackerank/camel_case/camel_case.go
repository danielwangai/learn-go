package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		print the number of words in a camel-cased word
	*/
	var input string
	numberOfWords := 1
	fmt.Println(strings.Split("string\none", "\n")[1])
	fmt.Scanf("%s\n", &input)
	for _, char := range input {
		/*
			char is the ASCII integer value for a character in the
			current stage of the loop
		*/
		charString := string(char)
		// compare uppercase value of current character
		if strings.ToUpper(charString) == charString {
			numberOfWords++
		}
	}
	fmt.Println(numberOfWords)
}
