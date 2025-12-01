package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("USER:", user)
	fmt.Println("HOME:", home)

	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}
	fmt.Println("Environment variable FRUIT set successfully.")
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	for _, e := range os.Environ() {
		kvpair := strings.SplitN(e, "=", 2)
		fmt.Println(kvpair[0])
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}
	fmt.Println("Environment variable FRUIT unset successfully.")
	fmt.Println("FRUIT env var after unset:", os.Getenv("FRUIT"))

	str := "a=b=c=d=e"
	// n = 1 returns "a=b=c=d"
	// n = 2 returns "a" and "b=c=d"
	// n = 3 returns "a" and "b" and "c=d"
	// n = 4 returns "a" and "b" and "c" and "d"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))

}