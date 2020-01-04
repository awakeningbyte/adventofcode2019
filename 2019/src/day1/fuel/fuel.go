package fuel

func Calculate(mass int) (int, error) {
	r := mass /3 - 2 
	if r < 0 {
		r = 0
	}
	
	return r, nil;
}

