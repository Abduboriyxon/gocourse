package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// current time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2022, time.December, 25, 10, 30, 0, 0, time.UTC)
	fmt.Println(specificTime)

	// parse time
	parsedTime, _ := time.Parse("2006-01-02", "2023-01-01") // Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "19-01-01") // Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "19-1-1") // Mon Jan 2 15:04:05 MST 2006
	parsedTime3, _ := time.Parse("06-1-2 15-04", "19-1-1 13-55") // Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// format time
	t := time.Now()
	fmt.Println(t.Format("Monday 2006-01-02 15:04:05"))

	onedayLater := t.Add(time.Hour * 24)
	fmt.Println(onedayLater)
	fmt.Println(onedayLater.Weekday())

	fmt.Println("Rounded Time:", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Tashkent")
	// t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	// // covert this to specific zone
	// tLocal := t.In(loc)

	// // perform rounding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original Time (UTC):", t)
	// fmt.Println("Original Time ( Local):", tLocal)
	// fmt.Println("Rounded Time (UTC):", roundedTime)
	// fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	fmt.Println("Truncate Time:", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")

	// covert time to location
	tInNY := time.Now().In(loc)
	fmt.Println("Time in New York:", tInNY)

	t1 := time.Date(2024, time.March, 10, 4, 30, 0, 0, time.UTC)
	t2 := time.Date(2024, time.March, 10, 1, 30, 0, 0, time.UTC)

	duration := t1.Sub(t2)
	fmt.Println("Duration between t1 and t2:", duration)

	// compare times
	fmt.Println("t1 before t2:", t1.After(t2))

}