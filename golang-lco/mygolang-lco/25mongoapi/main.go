package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rahuldevbodiga13/mongoapi/router"
)

func main() {
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":6000", router.InitRouter()))
	fmt.Println("Listening at port 6000...")

}
