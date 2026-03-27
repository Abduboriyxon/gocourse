package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by func(p1, p2 *Person) bool
}
func (s *personSorter) Len() int {
	return len(s.people)
}
func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}
func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}
func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	fmt.Println("Unsorting by age:", people)
	ageAsc := func (p1,p2 *Person) bool  {
		return p1.Age < p2.Age
	}
	name := func (p1,p2 *Person) bool  {
		return p1.Name < p2.Name
	}
	ageDesc := func (p1,p2 *Person) bool  {
		return p1.Age > p2.Age
	}
	lenName := func (p1,p2 *Person) bool  {
		return len(p1.Name) < len(p2.Name)
	}
	By(ageAsc).Sort(people)
	fmt.Println("Sorting by age (ascending):", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sorting by age (descending):", people)
	By(name).Sort(people)
	fmt.Println("Sorting by name:", people)
	By(lenName).Sort(people)
	fmt.Println("Sorting by length of name:", people)

	// ==== Example using sort.Slice ====
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorting strings by last character:", stringSlice)

	// ==== Example using sort.Interface directly ====
	// sort.Sort(ByAge(people))
	// numbers := []int{5, 2, 9, 1, 5, 6}
	// sort.Ints(numbers)
	// fmt.Println(numbers) // Output: [1 2 5 5 6 9]

	// stringSlice := []string{"banana", "apple", "cherry"}
	// sort.Strings(stringSlice)
	// fmt.Println(stringSlice) // Output: [apple banana cherry]
}


// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int{
// 	return len(a) 
// }
// func (a ByAge) Less(i, j int) bool{
// 	return a[i].Age < a[j].Age
// }
// func (a ByAge) Swap(i, j int){
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int{
// 	return len(a) 
// }
// func (a ByName) Less(i, j int) bool{
// 	return a[i].Name < a[j].Name
// }
// func (a ByName) Swap(i, j int){
// 	a[i], a[j] = a[j], a[i]
// }