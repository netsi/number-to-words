package main

import "testing"

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		want    string
		wantErr bool
	}{
		{name: "input below limit", num: 0, wantErr: true},
		{name: "input above limit", num: 10000, wantErr: true},
		{name: "input in limit", num: 1012, want: "one thousand and twelve"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NumberToWords(tt.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberToWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NumberToWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}
