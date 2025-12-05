package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name,omitempty"`
	Age int `xml:"age"`
	Address Address `xml:"address"`
	Email string `xml:"email"`
}

type Address struct {
	City string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	person := Person{Name: "John", Age: 30, Email: "email@examplemail.com", Address: Address{City: "New York", State: "NY"}}

	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatalln("Error marshalling to XML:", err)
	// }
	// fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling to XML:", err)
	}
	fmt.Println(string(xmlData1))

	// xmlRaw := `<person><name>Jane</name><age>25</age></person>`
	xmlRaw := `<person><name>John</name><age>25</age><address><city>San Francisco</city><state>CA</state></address></person>`
	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error unmarshalling from XML:", err)
	}
	fmt.Println(personxml)
	fmt.Println("local string",personxml.XMLName.Local)
	fmt.Println("namespace",personxml.XMLName.Space)

	book := Book{ISBN: "978-3-16-148410-0", Title: "Go Programming", Author: "Alice", Psudo: "psudo value", PsudoAttr: "psudo attr value"}
	xmlDataAttr, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling book to XML:", err)
	}
	fmt.Println(string(xmlDataAttr))
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN  string   `xml:"isbn,attr"`
	Title string   `xml:"title,attr"`
	Author string  `xml:"author,attr"`
	Psudo string  `xml:"psudo"`
	PsudoAttr string  `xml:"psudoattr,attr"`

}