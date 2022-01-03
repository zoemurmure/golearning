package geometry

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Triangle struct {
	size int
}

type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s *Square) SetSize(size float64) {
	s.size = size
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) SetRadius(r float64) {
	c.radius = r
}

type ColoredTriangle struct {
	Triangle
	color string
}

func (t Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t ColoredTriangle) Perimeter() int {
	return t.size * 3 * 2
}

func (t *Triangle) doubleSize() {
	t.size *= 2
}

//func main() {
//	t := triangle{6}
//	s := square{8}
//	t.doubleSize()
//
//	fmt.Println("Perimeter (triangle):", t.perimeter())
//	fmt.Println("Perimeter (square):", s.perimeter())
//
//	t2 := coloredTriangle{triangle{3}, "red"}
//	fmt.Println("Size:", t2.size)
//	fmt.Println("Perimeter(colored):", t2.perimeter())
//	fmt.Println("Perimeter(normal):", t2.triangle.perimeter())
//}
