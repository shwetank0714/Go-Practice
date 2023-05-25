package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){

	fmt.Println("Web requests Servers")
	//PerformGetRequest()
	PerformPostJsonRequest()
}


func PerformGetRequest()  {
	const myUrl = "http://localhost:8000/get"

	urlResponse, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}


	defer urlResponse.Body.Close()

	fmt.Println("Status code of Request", urlResponse.StatusCode)
	fmt.Println("Content length: ", urlResponse.ContentLength)
	dataByte , _ := ioutil.ReadAll(urlResponse.Body)
	fmt.Println(string(dataByte))


	// Another Method
	// Using strings package ---- 
	var responseString strings.Builder
	responseString.Write(dataByte)
	fmt.Println("New Method output: ", responseString.String())
}


func PerformPostJsonRequest()  {
	const posturl = "http://localhost:8000/post"

	// create a fake Json Payload

	requestBody := strings.NewReader(`
		{
			"name" : "shwetank",
			"age" : "23",
			"occupation" : "SDE",
			"Status" : "online"

		}
	`)

	postResponse, err := http.Post(posturl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer postResponse.Body.Close()

	urlByteFormat, _ := ioutil.ReadAll(postResponse.Body)

	fmt.Println("Post Data Sent: ", string(urlByteFormat))
}