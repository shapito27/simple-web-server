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
	Ru       string //create table word_translate. move there RU translation. in field language store ru
	ImgSrc   string
	Category int
	//Transcription string
	//Sound string
}

type Category struct {
	Id       int
	ParentId int
	NameEn   string
	NameRu   string
	ImgSrc   string
}

func getWords() (berries []Word) {
	strawberry := Word{1, "Strawberry", "Клубника", "http://www.calorizator.ru/sites/default/files/imagecache/product_192/product/strawberry-1.jpg", 2}
	gooseberry := Word{2, "Gooseberry", "Крыжовник", "https://www.veseys.com/media/catalog/product/cache/image/700x700/e9c3970ab036de70892d86c6d221abfe/3/7/37401-37401-image-37401-37401-image1-47368_2.jpg", 2}
	raspberry := Word{3, "Raspberry", "Малина", "https://james-mcintyre.co.uk/wp-content/uploads/2018/08/Raspberry_Glen-Lyon-500x500.jpg", 2}

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

	//fmt.Println(req.Method)
	//fmt.Println("path", req.URL.Path)

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
	checkError(err, "parse id is failed:")

	word, err := getWordById(id)
	checkError(err, "Getting word by id:")

	jsonRes, err := json.Marshal(word)
	checkError(err, "json err:")

	//response
	fmt.Fprintf(res, string(jsonRes))
}

func getCategories() []Category {
	categories := make([]Category, 3)

	categories[0] = Category{1, 0, "Food", "Еда", "/srctofood"}
	categories[1] = Category{2, 1, "Strawberry", "Ягоды", "/srctoberry"}
	categories[2] = Category{3, 1, "Nuts", "Орехи", "/srctonuts"}

	return categories
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "POST method not allowed")
	}

	categories := getCategories()
	res, err := json.Marshal(categories)
	checkError(err, "json err:")

	fmt.Fprintf(w, string(res))
}

func getCategoryById(id int) (Category, error) {
	categories := getCategories()

	for _, v := range categories {
		if v.Id == id {
			return v, nil
		}
	}

	return Category{}, errors.New("Category with id=" + string(id) + " not found")
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "POST method not allowed")
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkError(err, "parse id is failed:")

	cat, err := getCategoryById(id)
	checkError(err, "")

	res, err := json.Marshal(cat)
	checkError(err, "json err:")

	fmt.Fprintf(w, string(res))
}

func main() {
	//routing
	router := mux.NewRouter()
	router.HandleFunc("/", printInput)
	router.HandleFunc("/word", getAllWords)
	router.HandleFunc("/word/random", randomWord)
	router.HandleFunc("/word/{id}", getWord)
	router.HandleFunc("/category", getAllCategories)
	router.HandleFunc("/category/{id}", getCategory)
	http.Handle("/", router)

	err := http.ListenAndServe(":9091", nil) // set listen ip and port
	if err != nil {
		fmt.Println("\n\n", err)
	}
}

func checkError(err error, mes string) {
	if err != nil {
		fmt.Println(mes, err)
	}
}
