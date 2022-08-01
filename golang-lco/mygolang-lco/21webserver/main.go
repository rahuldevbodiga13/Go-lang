package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const base_url string = "http://localhost:8000/"

func main() {
	fmt.Println("Welcome to handling webserver req in Golang")
	// PerformGetRequest("get/")
	// PerformPostJsonRequest("post/")
	PerformPostFormRequest("postform/")
}

func PerformGetRequest(uri string) {
	reqUrl := base_url + uri
	res, err := http.Get(reqUrl)
	checkError(err)
	defer res.Body.Close()
	fmt.Println("Url is ", reqUrl)
	fmt.Println("Status code is ", res.StatusCode)
	fmt.Println("Content length is ", res.ContentLength)

	databytes, err := ioutil.ReadAll(res.Body)
	checkError(err)
	var responseString strings.Builder

	byteCount, _ := responseString.Write(databytes)

	fmt.Println("Byte count is ", byteCount)
	fmt.Println("Response is ", string(databytes))
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(uri string) {
	url := base_url + uri

	requestBody := strings.NewReader(`
		{
			"rahana" : "dev"
		}
		`)

	response, err := http.Post(url, "application/json", requestBody)
	checkError(err)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Response is ", string(databytes))

}

func PerformPostFormRequest(uri string) {
	myurl := base_url + uri

	data := url.Values{}
	data.Add("name", "dev")
	data.Add("age", "22")
	data.Add("desg", "tester/developer")

	response, err := http.PostForm(myurl, data)
	checkError(err)
	defer response.Body.Close()

	databytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Data is ", string(databytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
