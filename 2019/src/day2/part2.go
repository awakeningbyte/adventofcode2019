package main

import (
	"os"
	"fmt"
	"io/ioutil"
	. "day2/intcode"
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
	
	Reset(codes)
	Excute(codes)

	fmt.Printf("position 0 is %d \n", codes[0])


}
