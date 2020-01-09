package panel
import (
    "strconv"
    "fmt"
		"sort"
		"regexp"
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

func (p Panel) GetPoints() []Point {
	points := []Point{}
	for _,v := range p {
		points = append(points, v)
	}

	sort.SliceStable(points, func(i, j int) bool {
		return GetKey(points[i]) < GetKey(points[j])

	})
	return points
}

func (p Panel) Exist(pt Point) bool {
	key := GetKey(pt) 
	_, ok := p[key]
	return ok
}
func NewPanelFrom(input string) *Panel {
  regx,_ := regexp.Compile("[L|R|U|D]([1-9]+)")
	lines := regx.FindAllString(input, -1)
	var panel Panel= make(map[string]Point)
	start := Point{0,0}
	for _,l := range lines {
		points := start.Go(l)
		for _, p := range points {
			key := GetKey(p)
			panel[key]=p
		}
		start = points[len(points)-1]
	}
	return &panel
}

func abs(v int) int {
    if v < 0 {
        return v * -1
    }
    return v
}

func GetKey(pt Point) string {

	return fmt.Sprint("x",pt.X, "y", pt.Y)
}

