package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "hello, \ngo"
	message2 := "hello, \tgo"
	message3 := "hello, \rgo!"
	rawmessage := `hello, \ngo`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawmessage)

	fmt.Println(len(message))      // length in bytes
	fmt.Println(len(rawmessage))   // length in bytes

	fmt.Println("The first character is:", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple" // A has an ASCII value of 65
	str := "apple" // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app" // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, chac := range message {
		// fmt.Printf("Character at index %d is %c\n", i, chac)
		fmt.Printf("%v\n", chac)
	} 

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n%c\n", ch, jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type: %T Value: %v\n", cstr, cstr)

	const NIHONGO = "日本語" // Japanese language
	fmt.Println(NIHONGO)

	jhello := "こんにちば" // Hello in Japanese
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := "😀"
	fmt.Printf("%v\n", r)
}