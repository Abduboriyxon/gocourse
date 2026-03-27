package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	Age  int    `json:"age,omitempty"`
	Email string `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}

func main() {

	person := Person{FirstName: "John", Email: "sample@sample.com"}
	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Kane", Age: 25, Email: "fakemail.com", Address: Address{City: "New York", State: "NY"}}

	jsondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsondata1))

	jsonData1 := `{"full_name":"Jenny Doe", "emp_id": "009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`
	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Jenny's age increased by 5:",employeeFromJson.Age + 5)
	fmt.Println("Jenny's city:",employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "Los Angeles", State: "CA"},
		{City: "Chicago", State: "IL"},
		{City: "Houston", State: "TX"},
		{City: "Phoenix", State: "AZ"},
		{City: "Philadelphia", State: "PA"},
	}
	fmt.Println("List of City and State:", listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error marshalling list to JSON:", err)
	}
	fmt.Println("JSON List of City and State:", string(jsonList))

	// Handling unkown JSON structure
	jsonData2 := `{"name":"John", "age":30, "address":{"city":"New York", "state":"NY"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON to map:", err)
	}
	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
}

type Employee struct {
	FullName string `json:"full_name"`
	EmpID    string `json:"emp_id"`
	Age      int    `json:"age"`
	Address  Address `json:"address"`
}
