package main

import "testing"

func TestRomanNumeralValue(t *testing.T) {
	type args struct {
		RomanNumeral string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestMMXXII",
			args: args{RomanNumeral: "MMXXII"},
			want: 2022,
		},
		{
			name: "TestMMIX",
			args: args{RomanNumeral: "MMIX"},
			want: 2009,
		},
		{
			name: "TestDCLXVI",
			args: args{RomanNumeral: "DCLXVI"},
			want: 666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumeralValue(tt.args.RomanNumeral); got != tt.want {
				t.Errorf("RomanNumeralValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
