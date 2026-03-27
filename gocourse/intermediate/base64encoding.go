package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")

	// encode to base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// decode from base64
	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println(string(decode))

	// URL safe, avoids '+' and '/' characters
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL safe encoded:",urlSafeEncoded)
}