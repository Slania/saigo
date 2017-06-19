package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

////////////
// Square //
////////////

type Square struct {
	side float64
}

func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

////////////
// Circle //
////////////

type Circle struct {
	radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

////////////////
// Efficiency //
////////////////

type Shape interface {
	Name() string
	Perimeter() float64
	Area() float64
}

func Efficiency(s Shape) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

type FactoryShape struct {
	name      string
	area      float64
	perimeter float64
	Shape
}

func (factoryShape *FactoryShape) Name() string {
	return factoryShape.name
}

func (factoryShape *FactoryShape) Perimeter() float64 {
	return factoryShape.perimeter
}

func (factoryShape *FactoryShape) Area() float64 {
	return factoryShape.area
}

func Build(name string, args ...float64) func() Shape {
	return func() (shape Shape) {
		var factoryShape FactoryShape
		if args != nil {
			if len(args) > 1 {
				factoryShape = FactoryShape{name: name, area: args[0], perimeter: args[1]}
			} else {
				factoryShape = FactoryShape{name: name, area: args[0], perimeter: rand.Float64() * 10.0}
			}
		} else {
			factoryShape = FactoryShape{name: name, area: rand.Float64() * 10.0, perimeter: rand.Float64() * 10.0}
		}

		shape = &factoryShape
		return
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := Square{side: 10.0}
	Efficiency(&s)

	c := Circle{radius: 10.0}
	Efficiency(&c)

	generateFoo := Build("foo")
	for i := 1; i <= 10; i++ {
		Efficiency(generateFoo())
	}

	generateBar := Build("bar")
	for i := 1; i <= 10; i++ {
		Efficiency(generateBar())
	}
}
