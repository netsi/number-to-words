package main

import (
	"flag"
	"fmt"
	numberToWords "number-to-words/pkg/number-to-words"
	"os"
)

func main() {
	numPtr := flag.Int("num", 0, "an integer from 1 to 9999")
	flag.Parse()

	if numPtr == nil {
		fmt.Println("Please provide an integer from 1 to 9999 by using the -num flag")
		os.Exit(1)
	}

	words, err := NumberToWords(*numPtr)
	if err != nil {
		fmt.Printf("failed to convert the number to words with error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(words)
}

// NumberToWords receives a number and returns it in words accepts numbers from 1 to 9999
// e.g 	num=1 		returns one
//		num=12		returns twelve
//		num=56		returns fifty six
//		num=1012	returns one thousand and twelve
func NumberToWords(num int) (string, error) {
	if num < 1 || num > 9999 {
		return "", fmt.Errorf("input out of bounds: %d\n", num)
	}

	return numberToWords.NewThousandsNumber(num).ToWords()
}
