package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.dev:8080/lco?myName=shwetank&paymentid=90224ac34x"

func main()  {
	fmt.Println("Urls in Go lang")

	result,_ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	// this result.Quesry() returns the map of quesry params that are there in the url!!
	fmt.Println(result.Query())

	// IMPORTANT
	// CREATING URL FROM SOME DATA PRESENT
	// always pass the reference (&) while creating the url
	urlParts := &url.URL{
		Scheme: "https",
		Host: "punk.dev",
		Path: "/v1/live-tracing",
		RawPath: "user=Shwetank",
	}

	fmt.Println(urlParts.String())

}