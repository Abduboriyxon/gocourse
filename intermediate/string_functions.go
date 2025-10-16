package intermediate

import (
	"fmt"
	"strings"
)

func main() {

	// str := "Hello go!"

	// fmt.Println(len(str))          // 9

	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2 // Concatenation
	// fmt.Println(result)          // Hello World
	// fmt.Println(str[0])          
	// fmt.Println(str[1:7])

	// // String coversion
	// num := 42
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))
	
	// // string splitting
	// fruits := "apple, banana, cherry"
	// fruits1 := "apple-banana-cherry"
	// parts := strings.Split(fruits, ",")
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts) // [apple banana cherry]
	// fmt.Println(parts1)

	// countries := []string{"USA", "Canada", "Mexico"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined) // USA, Canada, Mexico

	// fmt.Println(strings.Contains(str, "go")) // true

	// replaced := strings.Replace(str, "go", "Universe", 1 )
	// fmt.Println(replaced) // Hello Universe!

	// strwspace := "   Hello World!   "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace)) // "Hello World!"
	// fmt.Println(strings.ToLower(str))          // hello go!
	// fmt.Println(strings.ToUpper(str))          // HELLO GO!

	// fmt.Println(strings.Repeat("Go! ", 3)) // Go! Go! Go!
	// fmt.Println(strings.Count("Hello", "l"))   // 2
	// fmt.Println(strings.HasPrefix("Hello", "He")) // true
	// fmt.Println(strings.HasSuffix("Hello", "lo")) // true

	// str5 := "Hel1lo, 123 Go 11!"
	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString(str5, -1)
	// fmt.Println(matches) // [1 123 11]

	// STRING BUILDER
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// convert builder to string
	result := builder.String()
	fmt.Println(result)

	// using writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("Go!")

	result = builder.String()
	fmt.Println(result)  

	// reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh")
	result = builder.String()
	fmt.Println(result)	

}