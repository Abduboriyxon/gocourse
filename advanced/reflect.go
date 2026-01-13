package advanced

import (
	"fmt"
	"reflect"
)

// ========== Basic Reflection ==========
// func main() {
// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Is Int:", t.Kind() == reflect.Int)
// 	fmt.Println("Is String:", t.Kind() == reflect.String)
// 	fmt.Println("Is zero:", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("v2 type:", v2.Type())
// 	fmt.Println("Original y:", v1.Int())

// 	v1.SetInt(18)
// 	fmt.Println("Modified y:", v1.Int())

// 	var itf interface{} = "hello"
// 	v3 := reflect.ValueOf(itf)
// 	fmt.Println("v3 type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("v3 string value:", v3.String())
// 	}
// }

// ========== Struct Field Reflection ==========
// type person struct{
// 	Name string
// 	age  int
// }
// func main() {
// 	p := person{Name: "Alice", age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField(){
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet(){
// 		nameField.SetString("Jane")
// 	}
// 	fmt.Println("Modified person:", p)
// }

// ========== Method Reflection ==========
type Greeter struct{}

func (g Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}
func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method
	fmt.Println("Type:", t)

	for i := range t.NumMethod(){
		method = t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	m := v.MethodByName("Greet")
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})
	fmt.Println("Greet Result:", results[0].String())
}
