package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {

	var r Rectangle
	var c Circle

	fmt.Print("Enter the co ordinates and radius for circle")
	fmt.Scan(&c.x, &c.y, &c.r)
	fmt.Print("Enter the co ordinates for Rectangle")
	fmt.Scan(&r.x1, &r.y1, &r.x2, &r.y2)
	fmt.Println("The circle co-ordinates and radius area :", c)
	fmt.Println("The rectangle co-ordinates are ", r)
	fmt.Println("Here's the output :")
	fmt.Println("****Area*******")
	fmt.Println("The area of circle: ", c.area())
	fmt.Println("The area of rectangle: ", r.area())
	fmt.Println("The total are of circle and Rectangle ", totalArea(&c, &r))
	fmt.Println("****Perimeter*****")
	fmt.Println("The perimeter of circle is :", c.perimeter())
	fmt.Println("The perimeter of rectangle is : ", r.perimeter())
	fmt.Println("The total perimeter of circle and Rectangle is :", totalPerimeter(&c, &r))
}
