package main

import (
    . "adventofcode2019/day3/panel"
    "io/ioutil"
    "os"
    "fmt"
)

type Result struct {
    Distance int
    CrossPoint *Point
}

func main() {
    text, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Errorf("can not open input file %s", os.Args[1])
    }

    part1 := Run(string(text))
    part2 := Run2(string(text))
    fmt.Printf("closest cross point is part1: %v, part2: %v\n", part1, part2)
}
