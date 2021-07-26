package number_to_words

import (
	"fmt"
	"strings"
)

// hundredsNumber stores the individual parts of a number up to the hundreds
type hundredsNumber struct {
	ones     int
	tens     int
	hundreds int
}

// NewHundredsNumber receives an integer and returns an instance to hundredsNumber
func NewHundredsNumber(num int) *hundredsNumber {
	return &hundredsNumber{
		ones:     num % 10,
		tens:     num / 10 % 10,
		hundreds: num / 100 % 10,
	}
}

// ToWords returns the number in words
// e.g 	10 => ten
// 		123 => one hundred and twenty three
func (n *hundredsNumber) ToWords() (string, error) {
	strBuilder := strings.Builder{}

	if n.hundreds != 0 {
		hundreds, err := getWordForBasicNumber(n.hundreds)
		if err != nil {
			return "", err
		}

		strBuilder.WriteString(fmt.Sprintf("%s %s", hundreds, Hundred))
	}

	if n.tens == 0 && n.ones == 0 {
		return strBuilder.String(), nil
	}

	tensWords, err := getWordsForTensAndOnes(n.tens, n.ones)
	if err != nil {
		return "", err
	}

	if strBuilder.Len() > 0 {
		strBuilder.WriteString(" ")
	}

	strBuilder.WriteString(tensWords)

	return strBuilder.String(), nil
}

// IsZero returns true if all the individual parts are zero
func (n *hundredsNumber) IsZero() bool {
	if n.ones != 0 {
		return false
	}

	if n.tens != 0 {
		return false
	}

	if n.hundreds != 0 {
		return false
	}

	return true
}
