package main

import "testing"

func Test(t *testing.T) {

	c := Circle{0, 1, 2}

	r := Rectangle{0, 2, 3, 4}

	cperi := c.perimeter()
	rperi := r.perimeter()

	if cperi != 12.566370614359172 {
		t.Error("This is not expected,got:", cperi)
	}

	if rperi != 10 {
		t.Error("This is not expected,got:", rperi)
	}
}
