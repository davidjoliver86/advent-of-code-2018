package main

import (
	"testing"
)

var parseLightPointData = []string{
	"position=< 9,  1> velocity=< 0,  2>",
	"position=< 7,  0> velocity=<-1,  0>",
	"position=< 3, -2> velocity=<-1,  1>",
	"position=< 6, 10> velocity=<-2, -1>",
}

var parseLightPointExpected = []*LightPoint{
	&LightPoint{Position: Point{9, 1}, Velocity: Point{0, 2}},
	&LightPoint{Position: Point{7, 0}, Velocity: Point{-1, 0}},
	&LightPoint{Position: Point{3, -2}, Velocity: Point{-1, 1}},
	&LightPoint{Position: Point{6, 10}, Velocity: Point{-2, -1}},
}

func TestMove(t *testing.T) {
	p := LightPoint{Position: Point{3, 9}, Velocity: Point{1, -2}}
	for i := 0; i < 3; i++ {
		p.Move()
	}
	if p.Position.X != 6 || p.Position.Y != 3 {
		t.Errorf("Test point expected (%d, %d), actual (%d, %d)", 6, 3, p.Position.X, p.Position.Y)
	}
}

func TestParseLightPoint(t *testing.T) {
	for i, source := range parseLightPointData {
		expected := parseLightPointExpected[i]
		lp := ParseLightPoint(source)
		pX := expected.Position.X == lp.Position.X
		pY := expected.Position.Y == lp.Position.Y
		vX := expected.Velocity.X == lp.Velocity.X
		vY := expected.Velocity.Y == lp.Velocity.Y
		if !(pX && pY && vX && vY) {
			t.Errorf("Test point expected (%d, %d), actual (%d, %d)", expected.Position.X, expected.Position.Y, lp.Position.X, lp.Position.Y)
		}
	}
}
