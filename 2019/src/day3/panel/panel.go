package panel
import (
    "strconv"
    "fmt"
)

type Point struct {
    X int
    Y int
}

func (p Point) Distance(o Point) int {
    return abs(p.X - o.X) + abs(p.Y - o.Y)
}

func (p Point) Go(line string) []Point {
    fmt.Printf("input line %s \n", line)
    n, _:= strconv.Atoi(line[1:])
    r := make([]Point,n,n)

    switch d :=line[:1]; d {
        case "L":
            for  j:=0; j<n; j++ {
                r[j]= Point{p.X - j -1, p.Y}
            }

        case "R":
            for  j:=0; j<n; j++ {
               r[j]= Point{p.X + j + 1, p.Y}
            }

        case "U":
            for  j:=0; j<n; j++ {
               r[j]= Point{p.X , p.Y + j + 1}
            }
        case "D":
            for  j:=0; j<n; j++ {
               r[j]= Point{p.X , p.Y - j -1}
           }
        default:
            panic("wrong line value")

    }

    return  r
}

func abs(v int) int {
    if v < 0 {
        return v * -1
    }
    return v
}

