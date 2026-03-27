package intermediate

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct{
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop()(T, bool){
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool{
	return len(s.elements)   == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack elements:")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {

	// x, y := 1, 2
	// x, y = swap(x, y)
	// println(x, y) // Output: 2 1

	// s1, s2 := "hello", "world"
	// s1, s2 = swap(s1, s2)
	// println(s1, s2) // Output: world hello

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty: ",intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty: ",intStack.isEmpty())
	
	stringStack := Stack[string]{}
	stringStack.push("a")
	stringStack.push("b")
	stringStack.push("c")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack empty: ",stringStack.isEmpty())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack empty: ",stringStack.isEmpty())

	
}