package strinter

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.00, 10.00}
	got := Perimeter(rect)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Rectnagle", func(t *testing.T) {
		rect := Rectangle{10.00, 10.00}
		checkArea(t, rect, 100.00)
	})

	t.Run("Circles", func(t *testing.T) {
		cir := Circle{10.00}
		checkArea(t, cir, 314.1592653589793)
	})

}
