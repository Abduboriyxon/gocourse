package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
	address Address
	HomePhoneCell
}

type Address struct{
	city string
	state string
}

type HomePhoneCell struct{
	home string
	cell string
}

func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) HaveBirthday() {
	p.age++
}

func main() {


	p1 := Person{
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "New York",
			state: "NY",
		},
		HomePhoneCell: HomePhoneCell{
			home: "123-456-7890",
			cell: "098-765-4321",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age: 25,
	}
	p2.address.city = "Los Angeles"
	p2.address.state = "USA "
	

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.FullName())
	fmt.Println(p2.address.city)
	fmt.Println(p1.address.city)
	fmt.Println(p1.HomePhoneCell.home)
	fmt.Println(p1.address.city )

	// Anonymous Struct
	user := struct {
		username string
		email string
	} {
		username: "johndoe",
		email: "pseudoemail@example.org",
	}

	fmt.Println(user.username, user.email)
	fmt.Println("Age before birthday:", p1.age)
	// p1.HaveBirthday()
	fmt.Println("Age after birthday:", p1.age)
	

}