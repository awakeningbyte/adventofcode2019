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

