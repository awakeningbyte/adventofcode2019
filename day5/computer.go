package computer

type Instruction struct {
    OpCode int
    Parameters []int
}
func Parse(code int) Instruction {
    c := code % 100
    p := code / 100
    switch c {
        case 1:
            params := getParams(p, 3) 
            return  Instruction{1, params} 
        case 2:
            params := getParams(p, 3) 
            return  Instruction{2, params} 
        case 3:
            return Instruction{3, {0}}
        case 4:
            return Insturction{4,{0}}
        default:
            panic("Unknow op code")
    }
}

func getParams(p int, l int) {

}
