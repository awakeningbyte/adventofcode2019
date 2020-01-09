package main

import (
    . "day3/panel"
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

    result := Run(string(text))
    fmt.Printf("closest cross point is %v\n", result)
}
