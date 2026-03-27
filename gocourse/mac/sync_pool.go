package mac

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age int
}

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new person")
			return &person{}
		},
	}

	// Get a object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 30
	fmt.Println("Got person from pool:", person1)

	fmt.Printf("Person1 - name: %s, age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person2:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got person3:", person3)

	// return the person to the pool
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Returned person2 and person3 to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got person4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person5:", person5)


}