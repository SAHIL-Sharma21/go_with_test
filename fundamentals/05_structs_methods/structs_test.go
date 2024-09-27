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

func TestArea(t *testing.T) {
	t.Run("area of rectnagle", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  20.0,
			Height: 40.0,
		}

		got := rectangle.Area()
		expected := 800.0

		if got != expected {
			t.Errorf("got %g and expected %g", got, expected)
		}
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10}

		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("got %g and expected %g", got, expected)
		}
	})
}
