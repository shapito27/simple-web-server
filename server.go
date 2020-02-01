package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Word struct {
	Id       int
	En       string
	Ru       string
	ImgSrc   string
	Category int
	//Transcription string
	//Sound string
}

func getWords() (berries []Word) {
	strawberry := Word{1, "Strawberry", "Клубника", "http://www.calorizator.ru/sites/default/files/imagecache/product_192/product/strawberry-1.jpg", 6}
	gooseberry := Word{2, "Gooseberry", "Крыжовник", "https://www.veseys.com/media/catalog/product/cache/image/700x700/e9c3970ab036de70892d86c6d221abfe/3/7/37401-37401-image-37401-37401-image1-47368_2.jpg", 3}
	raspberry := Word{3, "Raspberry", "Малина", "https://james-mcintyre.co.uk/wp-content/uploads/2018/08/Raspberry_Glen-Lyon-500x500.jpg", 9}

	berries = make([]Word, 0)
	berries = append(berries, strawberry)
	berries = append(berries, gooseberry)
	berries = append(berries, raspberry)

	return
}

func getWordById(id int) (Word, error) {
	words := getWords()
	for _, word := range words {
		if id == word.Id {
			return word, nil
		}
	}

	return Word{}, errors.New("Word with id =" + string(id) + " not found!")
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

func getWord(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	checkError("parse id is failed:", err)

	word, err := getWordById(id)
	checkError("Getting word by id:", err)

	jsonRes, err := json.Marshal(word)
	checkError("json err:", err)

	//response
	fmt.Fprintf(res, string(jsonRes))
}

func checkError(mes string, err error) {
	if err != nil {
		fmt.Println("json err:", err)
	}
}

func main() {
	//routing
	router := mux.NewRouter()
	router.HandleFunc("/", printInput)
	router.HandleFunc("/word/all", getAllWords)
	router.HandleFunc("/word/random", randomWord)
	router.HandleFunc("/word/{id}", getWord)
	http.Handle("/", router)

	err := http.ListenAndServe(":9091", nil) // set listen ip and port
	if err != nil {
		fmt.Println("\n\n", err)
	}
}
