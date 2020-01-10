package password

func CheckIncreasingRule(n int) bool {
	r := 10
	next  := n 
	for next > 0 {
		digit := next  % 10
		if digit > r {
			return false
		}
		r = digit
		next = next / 10
	}
	return true
}

func CheckDoubleRule(n int) bool {
	r := 10
	next  := n 
	for next > 0 {
		digit := next  % 10
		if digit == r {
			return true
		}
		r = digit
		next = next / 10
	}
	return false
}

func CheckLargerGroupRule(n int) bool {
	r := 10
	next  := n 
	flag := 0
	for next > 0 {
		digit := next  % 10
		if digit == r {
			flag += 1
		} else {
			if flag  == 1 {
				return true
			}
			flag =0
		}
		r = digit
		next = next / 10
	}
	return flag ==1
}


func Run(from , to int) (int, int) {
	p1 := 0
	p2 := 0
	for i:= from; i < to; i++ {
		if CheckIncreasingRule(i) && CheckDoubleRule(i)  {
			p1 +=1
			if CheckLargerGroupRule(i)  {
				p2 +=1
			}
		}
	}
	return p1,p2
}


