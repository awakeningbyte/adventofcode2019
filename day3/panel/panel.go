package panel
import (
    "strconv"
    "fmt"
		"sort"
		"regexp"
		"strings"
)

type Point struct {
    X int
    Y int
    Steps int
}

func (p Point) Distance(o Point) int {
    return abs(p.X - o.X) + abs(p.Y - o.Y)
}

func (p Point) Go(line string) []Point {
    n, _:= strconv.Atoi(line[1:])
    r := make([]Point,n,n)
    steps := p.Steps
    switch d :=line[:1]; d {
        case "L":
            for  j:=0; j<n; j++ {
                steps +=1
                r[j]= Point{p.X - j -1, p.Y, steps}
            }

        case "R":
            for  j:=0; j<n; j++ {
               steps +=1
               r[j]= Point{p.X + j + 1, p.Y, steps}
            }

        case "U":
            for  j:=0; j<n; j++ {
               steps +=1
               r[j]= Point{p.X , p.Y + j + 1, steps}
            }
        case "D":
            for  j:=0; j<n; j++ {
               steps +=1
               r[j]= Point{p.X , p.Y - j -1, steps}
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
		return points[i].Steps < points[j].Steps

	})
	return points
}

func (p Panel) Exist(pt Point) (Point, bool) {
	key := GetKey(pt) 
    o, f := p[key]
    return o, f
}
func NewPanelFrom(input string) *Panel {
  regx,_ := regexp.Compile("[L|R|U|D]([0-9]+)")
	lines := regx.FindAllString(input, -1)
	var panel Panel= make(map[string]Point)
	start := Point{0,0, 0}
	for _,l := range lines {
		points := start.Go(l)
		for _, p := range points {
			key := GetKey(p)
            _, ok := panel[key]
            if !ok {
			    panel[key]=p
            }
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

type CrossPoint struct {
	pt *Point
	dist int
}
func Run(wires string) int {
	found := CrossPoint{}
	lines := strings.Split(wires, "\n")
	panel := NewPanelFrom(lines[0])
	start := Point{0,0,0}
  regx,_ := regexp.Compile("[L|R|U|D]([0-9]+)")
	for _, l := range regx.FindAllString(lines[1], -1){
		points := start.Go(l)
		for _, p := range points {
            _, f :=panel.Exist(p)
			if f==true {
				d := p.Distance(Point{0,0,0})
				if found.pt == nil || d < found.dist {
					found = CrossPoint{&p, d}
				}
			}
		}
		start = points[len(points)-1]
	}
	return found.dist
}

func Run2(wires string) int {
	found := CrossPoint{}
	lines := strings.Split(wires, "\n")
	panel := NewPanelFrom(lines[0])
	start := Point{0,0,0}
    regx,_ := regexp.Compile("[L|R|U|D]([0-9]+)")
	for _, l := range regx.FindAllString(lines[1], -1){
		points := start.Go(l)
		for _, p := range points {
            o, f:= panel.Exist(p)
			if f == true {
				d := p.Steps + o.Steps
				if found.pt == nil || d < found.dist {
					found = CrossPoint{&p, d}
				}
			}
		}
		start = points[len(points)-1]
	}
	return found.dist
}


