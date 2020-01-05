package intcode_test 

import (
	"testing"
	"reflect"
	. "day2/intcode"
)

func TestReset(t *testing.T) {
	codes := []int{15,20, 22}
	codes,err := Reset(codes)
	if err != nil || codes[1] != 12 || codes[2] !=2 {
		t.Errorf("reset codes failed %v", codes)
	}
}

type excuteCase struct {
	codes []int
	expected []int
}

var excuteCases = []excuteCase{
	{[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
	{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
	{[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
	{[]int{2,4,4,5,99, 0}, []int{2,4,4,5,99,9801}},
	{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
}


func TestExcute(t *testing.T) {
	for _, test := range excuteCases {
		r, err := Excute(test.codes)
		if err != nil || !reflect.DeepEqual(test.expected, r) {
			t.Errorf("Expected %v but got %v", test.expected, r);
		}
	}
}

