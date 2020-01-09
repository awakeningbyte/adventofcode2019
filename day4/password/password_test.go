package password_test

import (
    "testing"
    . "adventofcode2019/day4/password"
)

type RuleTestCase struct  {
    num int
    valid bool
}

var ruleLengtTestCases = []RuleTestCase {
    {111111, true},
    {11111, false},
    {1111111, false},
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
