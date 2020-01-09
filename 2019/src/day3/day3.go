package main

import (
    . "day3/panel"
    "io/ioutil"
    "os"
    "strings"
    "regexp"
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

    result := Result{}
    var panel Panel 
    panel = *panel.Create()
    wires := strings.Split(string(text), "\n")
    
    regx,_:= regexp.Compile("[U|D|R|L]([1-9]+)")

    start := Point{0,0}
    for _, line :=  range regx.FindAllString(wires[0], -1) {
        points := start.Go(line)
        for _, p := range points {
            panel.Add(p)
        }
        start = points[len(points)-1]
    }

    for _, line :=  range regx.FindAllString(wires[1], -1) {
        points := start.Go(line)
        for _, p := range points {
            f := panel.Add(p)
            if f== true {
                dist := p.Distance(start)
                if result.CrossPoint == nil || result.Distance > dist {
                    result = Result{dist, &p}
                } 
            }
        }
        start = points[len(points)-1]
    }

    fmt.Printf("closest cross point is %v\n", result)
}
