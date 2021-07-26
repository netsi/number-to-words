package number_to_words

import (
	"fmt"
	"strings"
)

// thousandsNumber holds the information about a number up to the thousands
type thousandsNumber struct {
	hundreds  hundredsNumber
	thousands hundredsNumber
}

// NewThousandsNumber initializes thousandsNumber and returns a thousandsNumber instance with the given number
func NewThousandsNumber(num int) *thousandsNumber {
	hundreds := NewHundredsNumber(GetLastDigits(num, 3))
	thousands := NewHundredsNumber(GetLastDigits(num/1000, 3))

	return &thousandsNumber{
		hundreds:  *hundreds,
		thousands: *thousands,
	}
}

// ToWords returns the number in words if succeeds otherwise returns an error
func (n *thousandsNumber) ToWords() (string, error) {
	hundreds, err := n.hundreds.ToWords()
	if err != nil {
		return "", err
	}

	if n.thousands.IsZero() {
		return hundreds, nil
	}

	thousands, err := n.thousands.ToWords()
	if err != nil {
		return "", err
	}

	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("%s %s", thousands, Thousand))
	if !n.hundreds.IsZero() {
		strBuilder.WriteString(fmt.Sprintf(" %s %s", And, hundreds))
	}

	return strBuilder.String(), nil
}
