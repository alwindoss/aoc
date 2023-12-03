package day1

import "testing"

func Test_getCalibrationValFromString(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output int
	}{
		{
			desc:   "One",
			input:  "1abc2",
			output: 12,
		},
		{
			desc:   "Two",
			input:  "pqr3stu8vwx",
			output: 38,
		},
		{
			desc:   "Three",
			input:  "a1b2c3d4e5f",
			output: 15,
		},
		{
			desc:   "Four",
			input:  "treb7uchet",
			output: 77,
		},
		{
			desc:   "Fibr",
			input:  "pcg91vqrfpxxzzzoneightzt",
			output: 91,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := getCalibrationValFromString(tC.input)
			if s != tC.output {
				t.Fatalf("%s failed: expected %d but got %d", tC.desc, tC.output, s)
			}
		})
	}
}

func Test_run(t *testing.T) {

}
