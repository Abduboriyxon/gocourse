package basic

import (
	"fmt"
	foo "net/http"
)
func main(){
	fmt.Println("hello world")
	resp, err := foo.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
} 


