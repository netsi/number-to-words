package number_to_words_test

import (
	numberToWords "number-to-words/pkg/number-to-words"
	"testing"
)

func TestGetLastDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		x    uint
		want int
	}{
		{name: "single digit", num: 1, x: 10, want: 1},
		{name: "double digit", num: 24, x: 1, want: 4},
		{name: "four digit", num: 1024, x: 2, want: 24},
		{name: "eight digit", num: 10001024, x: 4, want: 1024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords.GetLastDigits(tt.num, tt.x); got != tt.want {
				t.Errorf("GetLastDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
