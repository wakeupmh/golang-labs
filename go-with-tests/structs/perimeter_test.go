package perimeter 

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expect float64) {
		t.Helper()

		result := shape.Area()

		if result != expect {
			t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
		}
	}

	t.Run("should return rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expect := 72.0

		verifyArea(t, rectangle, expect)
	})

	t.Run("should return circle area", func(t *testing.T) {
		circle := Circle{10.0}
		expect := 314.1592653589793

		verifyArea(t, circle, expect)
	})
}