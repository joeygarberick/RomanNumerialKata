package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p(RomanNumeralValue("IX"))
}

func convertValue(value string) int {
	switch value {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}

func RomanNumeralValue(RomanNumeral string) int {

	count := 0
	pos := 0
	values := s.Split(RomanNumeral, "")

	for pos < len(values)-1 {
		if convertValue(values[pos]) >= convertValue(values[pos+1]) {
			count += convertValue(values[pos])
		} else if convertValue(values[pos]) < convertValue(values[pos+1]) {
			count -= convertValue(values[pos])
		}
		pos++
	}
	count += convertValue(values[len(values)-1])

	return count
}
