package intermediate

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool("incorrect")

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)

	g := "hello こんにちは"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{255, 79, 65, 66, 227}
	j := string(i)
	fmt.Println(j)
}