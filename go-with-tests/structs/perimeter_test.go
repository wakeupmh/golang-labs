package perimeter 

import "testing"

func TestPerimeter(t *testing.T) {
	result := Perimeter(10.0, 10.0)
	expect := 40.0

	if result != expect {
		t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
	}
}

func TestArea(t *testing.T) {
    result := Area(12.0, 6.0)
    expect := 72.0

    if result != expect {
        t.Errorf("expect '%.2f' but got '%.2f'", expect, result)
    }
}