package main

import(
	"os"
	"strconv"
	"fmt"
	. "adventofcode2019/day4/password"
)

func main() {
	from,_ := strconv.Atoi(os.Args[1])
	to,_ := strconv.Atoi(os.Args[2])

	p1, p2 := Run(from, to)

	fmt.Printf("Part1 answer is %d, part2 answer is %d", p1, p2)
}
