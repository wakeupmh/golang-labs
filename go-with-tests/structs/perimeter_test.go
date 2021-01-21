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
	t.Run("should return rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := rectangle.Area()
		expect := 72.0

		if result != expect {
			t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
		}
	})

	t.Run("should return circle area", func(t *testing.T) {
		circle := Circle{10.0}
		result := circle.Area()
		expect := 314.1592653589793

		if result != expect {
			t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
		}
	})
}