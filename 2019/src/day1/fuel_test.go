package day1_test 

import (
	. "day1"
	"testing"
)
type calculationCase struct {
	mass int
	fuel int
	err error
}

var calculationTests = []calculationCase {
	{0,0, nil},
	{8,0, nil},
	{9,1,nil},
	{11, 1, nil },
	{12, 2, nil },
	{14, 2, nil },
	{1969, 654, nil },
	{100756, 33583, nil },
}
func TestCalculate(t *testing.T) {
	for _, test := range calculationTests {
		n, _:= Calculate(test.mass)
		if (n != test.fuel) {
			t.Errorf("Calculation result is incorrect. expected %d, actual %d", test.fuel, n)
		}

	}

}
