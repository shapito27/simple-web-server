package main

import (
	"fmt"
	"net/http"
)

func printInput(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	fmt.Println("Queries:")
	for k, v := range request.Form {
		fmt.Println(k, v)
	}
	fmt.Fprintf(response, "Hello Gopher!") // send data to client side
}

func main() {
	http.HandleFunc("/", printInput)         // set router
	err := http.ListenAndServe(":9091", nil) // set listen ip and port
	if err != nil {
		fmt.Println("\n\n", err)
	}
}
