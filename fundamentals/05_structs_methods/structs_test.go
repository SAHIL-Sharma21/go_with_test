package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f want %.2f", got, expected) //The f is for our float64 and the .2 means print 2 decimal places.
	}
}

// testing area
// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{Width: 20.0, Height: 40.0}
// 	got := Area(rectangle)
// 	expected := 800.0

// 	if got != expected {
// 		t.Errorf("got %.2f want %.2f", got, expected)
// 	}
// }

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g and want %g", got, want)
// 		}
// 	}

// 	t.Run("area of rectnagle", func(t *testing.T) {
// 		rectangle := Rectangle{
// 			Width:  20.0,
// 			Height: 40.0,
// 		}
// 		checkArea(t, rectangle, 800.0)
// 	})

// 	t.Run("area of circle", func(t *testing.T) {
// 		circle := Circle{Radius: 10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// refactoring thre test cases
// func TestArea(t *testing.T) {
// 	areaTest := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{12, 6}, 72.0},
// 		{Circle{10}, 314.1592653589793},
// 		{Triangle{12, 6}, 36.0},
// 	}

// 	for _, tt := range areaTest {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %g and want %g", got, tt.want)
// 		}
// 	}
// }

// refactoring abvove code for better developer experinece
func TestArea(t *testing.T) {

	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
