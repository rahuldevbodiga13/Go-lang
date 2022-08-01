package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=android&paymentid=sfdhfd232gbg"

func main() {
	fmt.Println("Handling Urls in golang")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.RequestURI())

	qparams := result.Query()
	fmt.Printf("Type of query params is : %T\n", qparams)
	fmt.Println(qparams["paymentid"])

	for key, value := range qparams {
		fmt.Println(key, value)
	}

	partsOfUri := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "tutcss/",
		RawPath: "user=rahul",
	}

	newUrl := partsOfUri.String()

	fmt.Println(newUrl)
}
