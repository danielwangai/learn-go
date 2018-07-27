package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// get user input
	inputWord := flag.String("word", "name", "Please enter a word/sentence to be converted to be encoded.")
	deltaValue := flag.Int("delta", 2, "Provide a number determining how to encode.")
	flag.Parse()

	if *deltaValue <= 0 {
		fmt.Println("This value has to be greater than zero.")
		return
	}

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/*
		convert inpput word to a slice of runes*/
	cypherRune := ""
	for _, v := range *inputWord {
		switch {
		case strings.IndexRune(alphabetUpper, v) >= 0:
			cypherRune = cypherRune + string(caesarCipher(v, *deltaValue, []rune(alphabetUpper)))
		case strings.IndexRune(alphabetLower, v) >= 0:
			cypherRune = cypherRune + string(caesarCipher(v, *deltaValue, []rune(alphabetLower)))
		default:
			cypherRune = cypherRune + string(v)
		}
	}

	printCipher(*inputWord, *deltaValue)
}

/*
Approach 1
*/
func caesarCipher(c rune, delta int, key []rune) rune {
	/*get the index of the rune in the key
	if character not found, returns -1
	*/
	index := strings.IndexRune(string(key), c)
	if index < 0 {
		panic("Error")
	}
	index = (index + delta) % len(key)
	return key[index]
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotateWithBase(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotateWithBase(r, 'a', delta)
	}
	// return original for complex ccharacters e.g. emojis, chinese alphabets
	return r
}

func rotateWithBase(r rune, base, delta int) rune {
	/*
		A=65, Z=90, a=97, z=122
	*/
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

// print out cyper
func printCipher(input string, delta int) {
	var cipherText []rune
	for _, v := range input {
		cipherText = append(cipherText, cipher(v, delta))
	}
	fmt.Println(string(cipherText))
}

/*
Approach 2
func caesarCipher(c rune, delta int, key []rune) rune {
	index := -1
	for i, v := range key {
		if v == c {
			index = i
			break
		}
	}
	if index < 0 {
		panic("Index < 0")
	}
	for i := 0; i < delta; i++ {
		index++
		if index >= len(key) {
			index = 0
		}
	}
	return key[index]
}
*/
