package task2

import (
	"fmt"
	"math"
)

//面向对象

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r

}

func Object() {
	r := Rectangle{2, 3}
	fmt.Printf("Rectangle r.Area(): %v\n", r.Area())
	fmt.Printf("Rectangle r.Perimeter(): %v\n", r.Perimeter())

	c := Circle{3}
	fmt.Printf("Circle c.Area(): %v\n", c.Area())
	fmt.Printf("Circle c.Perimeter(): %v\n", c.Perimeter())

	e := Employee{Person{"zhangsan", 20}, 1}
	e.PrintInfo()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	person     Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工: %v\n", e)
}
