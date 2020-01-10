package password_test

import (
    "testing"
    . "adventofcode2019/day4/password"
)

type RuleTestCase struct  {
    num int
    valid bool
}

var ruleIncreasingCases = []RuleTestCase {
    {123456, true},
    {223450, false},
    {223456, true},
}

var ruleDoubleCases = []RuleTestCase {
    {123489, false},
    {122489, true},
}

func TestIncreaseRule(t *testing.T) {
	for _, test := range ruleIncreasingCases {
		ok := CheckIncreasingRule(test.num)
		if ok != test.valid {
			t.Errorf("check %v for increasing rule, expect %v, got %v", test.num, test.valid, ok) 
		}
	}
}

func TestDoubleRule(t *testing.T) {
	for _, test := range ruleDoubleCases {
		ok := CheckDoubleRule(test.num)
		if ok != test.valid {
			t.Errorf("check %v for double digits rule, expect %v, got %v", test.num, test.valid, ok) 
		}
	}
}
