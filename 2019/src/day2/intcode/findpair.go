package intcode
func Tryfind(output int, origcodes []int) (noun int, verb int) {
	codes := make([]int,len(origcodes), len(origcodes))
	for x  := 0; x < len(codes); x++ {
		for y :=0; y < len(codes); y++ {
			copy(codes, origcodes)
			codes, _ = Reset(codes, x, y)			
			c, err := Excute(codes)
			if err == nil && c[0] == output {
				return x, y
			}
		}
	}
	return -1, -1
}

