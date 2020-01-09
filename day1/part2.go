package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	. "adventofcode2019/day1/fuel"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Can not open %s\n",os.Args[1])
	}
	defer file.Close()
	
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Can not parse string %s to interger\n", line)
			return
		}

		n, err := Calculate2(mass)
				
		if err != nil {
			fmt.Printf("There is an error while calculating %d", mass)
			return
		}
		total += n

	}
	
	fmt.Printf("Fuel in total %d", total)
}

