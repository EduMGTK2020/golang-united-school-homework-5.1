package square

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {

	startPoint := Point{x: 10, y: 5}
	testSquare := Square{start: startPoint, a: 5}

	end := testSquare.End()
	expEnd := Point{x: 15, y: 10}
	if !reflect.DeepEqual(end, expEnd) {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expEnd, end)
	}

	area := testSquare.Area()
	expArea := uint(25)
	if area != expArea {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expArea, area)
	}

	perimeter := testSquare.Perimeter()
	expPerimeter := uint(20)
	if perimeter != expPerimeter {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expPerimeter, perimeter)
	}

}
