package number_to_words_test

import (
	numberToWords "number-to-words/pkg/number-to-words"
	"testing"
)

func Test_thousandsNumber_ToWords(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		want    string
		wantErr bool
	}{
		{name: "single digit", num: 1, want: "one", wantErr: false},
		{name: "two digits", num: 12, want: "twelve", wantErr: false},
		{name: "two digits > 19", num: 56, want: "fifty six", wantErr: false},
		{name: "two digits > 19 round number", num: 50, want: "fifty", wantErr: false},
		{name: "four digits", num: 1012, want: "one thousand and twelve", wantErr: false},
		{name: "three digits", num: 212, want: "two hundred twelve", wantErr: false},
		{name: "four digits with hundreds", num: 3212, want: "three thousand and two hundred twelve", wantErr: false},
		{name: "max supported number", num: 9999, want: "nine thousand and nine hundred ninety nine", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := numberToWords.NewThousandsNumber(tt.num)
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
