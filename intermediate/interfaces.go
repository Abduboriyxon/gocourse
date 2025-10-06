package intermediate

import (
	"fmt"
	"math"
)

type gemotry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}


func (r rect) area() float64{
	return r.height * r.width
}

func (r rect) perim() float64{
	return 2 * (r.height + r.width)
}


type circle struct {
	radius float64
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64{
	return 2 * c.radius
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius 
}

type rect1 struct {
	width, height float64
}

func (r rect1) perim() float64{
	return 2 * (r.height + r.width)
}

func (r rect1) area() float64{
	return r.height * r.width
}

func measure(g gemotry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// r1 := rect1{width: 3, height: 4}


	measure(r)
	measure(c)
	// measure(r1)

	myPrinter(1, "hello", 3.14, r, c)
	printType(1)
	printType("hello")
	printType(3.14)

}

func printType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("Type is int\n")
	case string:
		fmt.Printf("Type is string\n")
	default:
		fmt.Printf("Type is unknown\n")
	}
}

func myPrinter (i ...interface{}){
	for _, v := range i {
		fmt.Println(v)
	}
}