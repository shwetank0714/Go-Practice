package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"
func main()  {
	
	fmt.Println("Web Requests")

	// response objects gets (*http.response) as return
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// responseDataBytes contain the response body
	responseDataBytes, _ := ioutil.ReadAll(response.Body)

	// converts the byte array into a readable string format ( Here in this example it is html)
	responseBody := string(responseDataBytes)
	fmt.Println("The response Body is: \n", responseBody)
}