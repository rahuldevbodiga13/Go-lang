package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"cource_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling json in golang")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs BootCamp", 100, "Lco.in", "abc123", []string{"web-dev", "js"}},
		{"Mern BootCamp", 299, "Lco.in", "bcd123", nil},
		{"Android BootCamp", 1999, "Lco.in", "fds123", []string{"mobile-dev", "js"}},
	}

	//package the data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonResponse := []byte(`
	{
		"cource_name": "ReactJs BootCamp",
		"price": 100,
		"website": "Lco.in",
		"tags": ["web-dev","js"]
	}
	`)

	var decodedData course

	if json.Valid(jsonResponse) {
		err := json.Unmarshal(jsonResponse, &decodedData)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", decodedData)
		fmt.Println("Decoded data is \n", decodedData)
	} else {
		fmt.Println("Json was not valid")
	}

	//there are some cases we dont require structs, instead we make use of key,value pairs

	var mappedData map[string]interface{}
	err := json.Unmarshal(jsonResponse, &mappedData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", mappedData)

}
