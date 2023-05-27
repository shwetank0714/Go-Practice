package main

import (
	"encoding/json"
	"fmt"
)

// Alias for the name
// this "-" depicts that in the response the data corresponding to that field will not be shown in json
type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"courseprice"`
	Platform string
	Password string `json:"-"`
	Tags     []string
}

func main() {
	fmt.Println("welcome to Jsons")
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{
			Name:     "Shwetank",
			Price:    234,
			Platform: "VsCode",
			Password: "kpmgDetug9099",
			Tags:     []string{"hello", "world", "how", "are", "you"},
		},
		{
			Name:     "John",
			Price:    290,
			Platform: "Ecllipse",
			Password: "kpmgDetug9099",
			Tags:     []string{"hello", "world", "how", "are", "you"},
		},
		{
			Name:     "Roman",
			Price:    634,
			Platform: "jetBrains",
			Password: "kpmgDetug9099",
			Tags:     []string{"hello", "world", "how", "are", "you"},
		},
		{
			Name:     "Shwetank",
			Price:    234,
			Platform: "VsCode",
			Password: "kpmgDetug9099",
			Tags:     []string{"hello", "world", "how", "are", "you"},
		},
	}

	_, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}

	indentedJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(indentedJson))
}
