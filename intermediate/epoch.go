package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// 00:00:00 UTC on 1 January 1970

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix time:", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println("Converted back to time:", t)
	fmt.Println("Formatted time:", t.Format("2006-01-02"))


}