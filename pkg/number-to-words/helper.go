package number_to_words

import (
	"errors"
	"fmt"
	"math"
)

// GetLastDigits returns the last X digits from num
// e.g GetLastDigits(1024, 2) returns 24
func GetLastDigits(num int, x uint) int {
	return num % int(math.Pow(float64(10), float64(x)))
}

// getWordForBasicNumber returns the word for a basic number, basic numbers are considered 1~19
func getWordForBasicNumber(num int) (word string, err error) {
	word, ok := basicNums[num]
	if !ok {
		return "", errors.New("basic number not found")
	}

	return
}

// getWordsForTensAndOnes returns the words for tens and ones (numbers 1 to 99)
func getWordsForTensAndOnes(tens int, ones int) (words string, err error) {
	if tens == 1 || tens == 0 {
		return getWordForBasicNumber(tens*10 + ones)
	}

	// is a round number
	if ones == 0 {
		return getWordForTens(tens * 10)
	}

	tensWord, err := getWordForTens(tens * 10)
	if err != nil {
		return
	}

	onesWord, err := getWordForBasicNumber(ones)
	if err != nil {
		return
	}

	return fmt.Sprintf("%s %s", tensWord, onesWord), nil
}

// getWordForTens return word for tens
func getWordForTens(num int) (word string, err error) {
	word, ok := tens[num]
	if !ok {
		return "", errors.New("tens not found")
	}

	return
}
