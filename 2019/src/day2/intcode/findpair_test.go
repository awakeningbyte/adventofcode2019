package intcode_test

import (
	"testing"
	. "day2/intcode"
)

type findpairCase struct {
	codes []int
	output int
}

var findpairCases = []findpairCase{
	{[]int{1,9,10,3,2,3,11,0,99,30,40,50},3500},
	{[]int{1,0,0,0,99},2},
	{[]int{2,3,0,3,99},2},
	{[]int{2,4,4,5,99, 0},2},
	{[]int{1,1,1,4,99,5,6,0,99},30},
}

func TestTryfind(t *testing.T) {
	for _, test := range findpairCases {
		noun, verb := Tryfind(test.output, test.codes )
		input, _ := Reset(test.codes, noun, verb)
		c,err := 	Excute(input)
		if err != nil || c[0] != test.output {
			t.Errorf("Test failed, %v, nount %d, verb %d", test, noun, verb)
		}
	}

}


