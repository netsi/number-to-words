package number_to_words_test

import (
	numberToWords "number-to-words/pkg/number-to-words"
	"testing"
)

func Test_hundredsNumber_ToWords(t *testing.T) {
	tests := []struct {
		name     string
		inputNum int
		want     string
		wantErr  bool
	}{
		{name: "one digit", inputNum: 1, want: "one"},
		{name: "two digits < 20", inputNum: 19, want: "nineteen"},
		{name: "two digits >= 20 ", inputNum: 99, want: "ninety nine"},
		{name: "three digits", inputNum: 108, want: "one hundred eight"},
		{name: "three digits round", inputNum: 100, want: "one hundred"},
		{name: "three digits no ones", inputNum: 110, want: "one hundred ten"},
		{name: "max three digits", inputNum: 999, want: "nine hundred ninety nine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := numberToWords.NewHundredsNumber(tt.inputNum)
			got, err := n.ToWords()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}
