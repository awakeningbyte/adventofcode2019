package password_test

import (
    "testing"
    . "adventofcode2019/day4/password"
)


var ruleLargerGroupCases = []RuleTestCase {
    {112233, true},
    {123444, false},
    {111244, true},
    {124444, false},
    {123456, false},
}

func TestLargerGroupRule(t *testing.T) {
	for _, test := range ruleLargerGroupCases {
		ok := CheckLargerGroupRule(test.num)
		if ok != test.valid {
			t.Errorf("check %v for larger group rule, expect %v, got %v", test.num, test.valid, ok) 
		}
	}
}
