package intcode 
func Reset(codes []int) ([]int ,error) {
	codes[1] = 12
	codes[2] = 2
	return codes, nil
}

func Excute(codes []int) ([]int ,error) {
	for i := range codes {
		if i % 4 == 0  {
			c := codes[i]
			switch c {
				case 99 :
					return codes, nil
				case 1:
					codes[codes[i+3]] = codes[codes[i+1]]+codes[codes[i+2]]
				case 2:
					codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
			}
		}
	}
	return codes, nil
}
