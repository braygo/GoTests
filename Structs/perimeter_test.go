package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 2.0}
	got := Perimeter(rectangle)
	want := 28.0

	if want != got {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 2.0}
		got := rectangle.Area()
		want := 24.0

		if want != got {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := math.Pi * 100

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
