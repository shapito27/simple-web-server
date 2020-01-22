package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"os"
	"math/rand"
	"time"
)

type Word struct {
	En       string
	Ru       string
	Src      string
	Category int
}

func getWords() (berries []Word) {
	strawberry := Word{"Strawberry", "Клубника", "http://www.calorizator.ru/sites/default/files/imagecache/product_192/product/strawberry-1.jpg", 6}
	gooseberry := Word{"Gooseberry", "Крыжовник", "https://www.veseys.com/media/catalog/product/cache/image/700x700/e9c3970ab036de70892d86c6d221abfe/3/7/37401-37401-image-37401-37401-image1-47368_2.jpg", 3}
	raspberry := Word{"Raspberry", "Малина", "https://james-mcintyre.co.uk/wp-content/uploads/2018/08/Raspberry_Glen-Lyon-500x500.jpg", 9}

	berries = make([]Word, 0)
	berries = append(berries, strawberry)
	berries = append(berries, gooseberry)
	berries = append(berries, raspberry)

	return
}

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

func getAllWords(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintf(res, "POST method not allowed")
	}

	arr := getWords()
	jsonRes, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(req.Method)
	fmt.Println("path", req.URL.Path)

	//response
	fmt.Fprintf(res, string(jsonRes))
}

func randomWord(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintf(res, "POST method not allowed")
	}

	arr := getWords()
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	word := arr[rand.Intn(len(arr)-1)]
	jsonRes, err := json.Marshal(word)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(req.Method)
	fmt.Println("path", req.URL.Path)

	//response
	fmt.Fprintf(res, string(jsonRes))
}

func main() {
	//routing
	http.HandleFunc("/", printInput)
	http.HandleFunc("/trainig/all-words", getAllWords)
	http.HandleFunc("/trainig/random-word", randomWord)

	err := http.ListenAndServe(":9091", nil) // set listen ip and port
	if err != nil {
		fmt.Println("\n\n", err)
	}
}
