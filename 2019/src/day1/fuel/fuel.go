package fuel

func Calculate(mass int) (int, error) {
	r := mass /3 - 2 
	if r < 0 {
		r = 0
	}
	
	return r, nil;
}

func Calculate2(mass int) (int, error) {
	f :=0
	r,err := Calculate(mass)
	for r >0 && err == nil {
		f += r
		r, err =Calculate(r)
	}


	return f, nil
}

