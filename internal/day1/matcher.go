package day1

import "strings"

var digits = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

type Data struct {
	Input       string
	DigitsFound []struct {
		Digit string
		Index int
	}
}

func getFirstAndLastDigitConcatenated(in string) {
	d := Data{}
	for i := 0; i < len(digits)-1; i++ {
		d.Input = digits[i]
		if strings.Contains(in, digits[i]) {
			index := strings.Index(in, digits[i])
			d.DigitsFound = append(d.DigitsFound, struct {
				Digit string
				Index int
			}{digits[i], index})
		}
	}
}
