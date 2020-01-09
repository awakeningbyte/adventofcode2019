package panel_test

import (
    . "day3/panel"
    "reflect"
    "testing"
		"sort"
)

type PointsDistanceTestCase struct {
    From Point
    To Point
    Distance int
}

var distanceTestCases = []PointsDistanceTestCase{
    { Point {0, 0 }, Point {0, 0 }, 0},
    { Point {0, 1 }, Point {1, 0 }, 2},
    { Point {1, 1 }, Point {0, 0 }, 2},
    { Point {4, 3 }, Point {1, 10 }, 10},
}

type LineParsingTestCase struct {
    start Point
    line string
    points []Point
}

var lineParsingTestCases = []LineParsingTestCase {
    {Point {0, 0}, "R2", []Point{Point{1,0}, Point{2,0}}},
    {Point {0, 0}, "U1", []Point{Point{0,1}}},
    {Point {0, 1}, "D1", []Point{Point{0,0}}},
    {Point {2, 2}, "L1", []Point{Point{1,2}}},
    {Point {5, 5}, "R3", []Point{Point{6,5}, Point{7,5}, Point{8,5}}},
    {Point {2, 2}, "L1", []Point{Point{1,2}}},
    {Point {5, 2}, "U2", []Point{Point{5,3}, Point{5,4}}},
}


func TestPointsDistanceCalculation(t *testing.T) {
    for _, test := range distanceTestCases {
        d := test.From.Distance(test.To)
        if d != test.Distance {
            t.Errorf("Wrong calcualtion of distance from %v to %v, expected %d, got %d", test.From, test.To, test.Distance, d) 
        }
    }
}

func TestLineParsing(t *testing.T) {
    for _, test :=range  lineParsingTestCases {
        points := test.start.Go(test.line)
        if !reflect.DeepEqual(points, test.points) {
            t.Errorf("Parsing line returns incorrect result. From %v go %s, expect %v, got %v", test.start, test.line, test.points, points)
        }
    }
}

type PointExistenceTestCase struct {
    wires string
    point Point
    exist bool
}

var pointExistenceTestCases = []PointExistenceTestCase {
	{"R2,U2,L2,D1", Point{2,2}, true},
	{"R2,U2,L2,D1", Point{3,0}, false},
	{"R2,U2,L2,D1" ,Point{0,1}, true},
	{"R2,U2,L2,D1,L3,D1\n", Point{-3,0}, true},
	{"R2,U2,L2,D1", Point{3,0}, false},
}

func TestPointExist(t *testing.T) {	
	for _, test := range pointExistenceTestCases {
		panel := NewPanelFrom(test.wires)
		actual := panel.Exist(test.point)

		if actual!= test.exist {
			t.Errorf("point %v existence in %v should be %v but got %v", test.point,panel.GetPoints(), test.exist, actual)
		}
	}
}

func TestPanelConstruction(t *testing.T) {
	panel:=NewPanelFrom("R2,U2,L2,D1")
	expected := []Point{ Point{1,0}, Point{2,0}, Point{2,1}, Point{2,2}, Point{1,2}, Point{0,2}, Point{0,1}}
	actual := panel.GetPoints()

	sort.SliceStable(expected, func(i, j int) bool {
		return GetKey(expected[i]) < GetKey(expected[j])
	})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("panel construction wrong, expected %v, got %v", expected, actual)
	}
}

type WiresTestCase struct {
	wires string
	distance int
}

var wiresTestCases = []WiresTestCase {
	{"R8,U5,L5,D3\nU7,R6,D4,L4\n", 6},
	{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83\n",159},
	{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",135},
}

func TestRun(t *testing.T) {
	for _,test :=range wiresTestCases {
		d := Run(test.wires)
		if d!=test.distance {
			t.Errorf("Run returns wrong distance, expect %d, got %d", test.distance, d)
		}
	}
}


