package main

import (
	"os"
	"fmt"
	"io/ioutil"
	. "adventofcode2019/day2/intcode"
	"strconv"
	"strings"
)

func main() {
	text , err:=  ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("can not open file %s", os.Args[1])
	}
	str := strings.TrimRight(string(text), "\n")
	arr := strings.Split(str, ",")
	codes := make([]int, 0, len(arr))
	
	for _, s := range arr {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		codes = append(codes, n)

	}
	
	noun, verb := Tryfind(19690720, codes)
	fmt.Printf("part2 anwser is %d \n", noun * 100 + verb)


}
