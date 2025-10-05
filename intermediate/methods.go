package intermediate

import "fmt"

type Regtangular struct{
	length float64
	width float64
}

type Shape struct{
	Regtangular
}

// method with value reciver
func (r Regtangular) Area() float64 {
	return r.length * r.width
}

// method with pointer reciver
func (r *Regtangular) Scale (factor float64) {
	r.length *= factor
	r.width *= factor 
}

func main() {
	rect := Regtangular{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area wirh rectangular: ",area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area wirh rectangular: ",area)

	num := MyInt(-5)
	num1 := MyInt(10)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage()) 
	
	s := Shape{Regtangular{length: 5, width: 4}}
	fmt.Println("Area with embedded struct: ", s.Area())


}

type MyInt int

// method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string{
	return "Welcome to Go methods!"
}