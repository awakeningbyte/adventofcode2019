package computer_test
import (
    'testing'
    . 'computer'
)
type ParseInstructionTestCase struct {
    opt int
    instr Instruction
}

var parseInstructionTestCases = ParseInstructionTestCase {
    {3, {3, {0}}} ,
    {4, {4, {0}}},
    {1101, {1, {1,1,0}}},
}

func TestPasre(t *testing.T) {
    for _, test := range paseInstructionTestCases {
        instr:= Parse(test.opt)
        if instr != test.instr {
            t.Error("Parsing input incorrect, expecting %v, got %v", test.instr, instr);
        }
    }
}
