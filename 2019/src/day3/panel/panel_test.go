package panel_test

import (
    . "day3/panel"
    "reflect"
    "testing"
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
