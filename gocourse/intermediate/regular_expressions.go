package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"Hello, World!\"")
	fmt.Println(`He said, "Hello, World!"`)

	// compile a regex pattern to macth eamil address
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// test strings
	email1 := "user@email.com"
	email2 := "invalid-email"

	// match
	fmt.Println("Email 1 valid: ", re.MatchString(email1))
	fmt.Println("Email 2 valid: ", re.MatchString(email2))

	// capturing groups
	// compile regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// test date string
	date := "2024-06-15"

	// find submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// target string
	str := "Hello world"

	re = regexp.MustCompile(`[aeoiu]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive matching
	// m - multi-line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	// test strings
	text := "Go is an open source programming language. I love GO!"

	// match
	fmt.Println("Match:", re.MatchString(text))


}