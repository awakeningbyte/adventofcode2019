package computer

func Parse(code int) {
    c := code % 100
    p := code / 100
    switch c {
        case 1:
            return  Instruction{1, {}} 
        case 2:
        case 3:
        case 4:
        default:
            panic("Unknow op code")
    }


}
