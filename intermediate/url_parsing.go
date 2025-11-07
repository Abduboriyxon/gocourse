package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURl1 := "https://example.com/path?name=John&age=30"

	parseURl1, err := url.Parse(rawURl1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	queryParams := parseURl1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:  "example.com",
		Path: "/path",
	}
	query := baseURL.Query()
	query.Set("name", "Alice")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()
	fmt.Println("Built URL:", baseURL.String())

	values := url.Values{}

	// Add key-value pairs to the values object
	values.Add("name", "Kevin")
	values.Add("age", "25")
	values.Add("city", "NewYork")
	values.Add("country", "USA")

	// Encode the values to a URL-encoded query string
	encodeQuery := values.Encode()
	fmt.Println("Encoded Query String:", encodeQuery)

	// Build url
	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodeQuery
	fmt.Println(fullURL)
}