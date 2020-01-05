package fuel_test 

import (
	. "day1/fuel"
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
var calculationTests2 = []calculationCase {
	{14,2, nil},
	{1969,966, nil},
	{100756,50346,nil},
}
func TestCalculate(t *testing.T) {
	for _, test := range calculationTests {
		n, _:= Calculate(test.mass)
		if (n != test.fuel) {
			t.Errorf("Calculation result is incorrect. expected %d, actual %d", test.fuel, n)
		}

	}

}

func TestCalculate2(t *testing.T) {
	for _, test := range calculationTests2 {
		n, _:= Calculate2(test.mass)
		if (n != test.fuel) {
			t.Errorf("Calculation2 result is incorrect. expected %d, actual %d", test.fuel, n)
		}

	}

}
