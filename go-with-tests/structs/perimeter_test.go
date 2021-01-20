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
	rectangle := Rectangle{12.0, 6.0}
    result := Area(rectangle)
    expect := 72.0

    if result != expect {
        t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
    }
}