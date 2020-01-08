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

type Panel map[string]Point 

func (p Panel) Create() *Panel {
    var n Panel =make(map[string]Point)
    return &n
}

func (p Panel) Add(pt Point) bool {
    key := fmt.Sprint("X", pt.X, "Y", pt.Y)
    _, ok :=  p[key]
    if ok {
        return true
    }
    p[key] = pt
    return false
}

func abs(v int) int {
    if v < 0 {
        return v * -1
    }
    return v
}

