package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "postgres://user:pass@host.com:5432/path?k=v#f"
	parsed, err := url.Parse(urlString)
	if err != nil {
		panic("oh no")
	}

	fmt.Println(parsed.Scheme)
	fmt.Println(parsed.Opaque)
	fmt.Println(parsed.User)
	fmt.Println(parsed.Port())
	fmt.Println(parsed.Path)
	fmt.Println(parsed.RawQuery)
	fmt.Println(parsed.Fragment)
}
