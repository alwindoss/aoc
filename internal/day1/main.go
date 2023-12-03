package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	input := "./data/day1_input.txt"
	output := "./data/day1_output.txt"

	var listOfCalibrationVals []int
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	scnr := bufio.NewScanner(f)
	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	for scnr.Scan() {
		line := scnr.Text()
		sum := getCalibrationValFromString(line)
		out.WriteString(fmt.Sprintf("%d\n", sum))
		listOfCalibrationVals = append(listOfCalibrationVals, sum)
	}
	var total int
	for _, v := range listOfCalibrationVals {
		total = total + v
	}
	fmt.Println("The total calibration values is:", total)
}

func getCalibrationValFromString(in string) int {
	var digits []string
	for _, s := range in {
		if '0' <= s && s <= '9' {
			str, err := strconv.Unquote(strconv.QuoteRune(s))
			if err != nil {
				panic(err)
			}
			digits = append(digits, str)
		}
	}
	l := len(digits)
	v := digits[0] + digits[l-1]
	re, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return re
}
