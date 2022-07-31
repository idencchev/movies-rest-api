package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var movies []interface{}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://api.tvmaze.com/shows")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &movies)
	json.NewEncoder(w).Encode(movies)
}

func SearchMovies(w http.ResponseWriter, r *http.Request) {
	var url string

	query := strings.Split(r.URL.RawQuery, "=")

	if query[0] == "movie" {
		fmt.Println(query[0])
		url = fmt.Sprintf("https://api.tvmaze.com/search/shows?q=%s", query[1])
	} else if query[0] == "title" {
		fmt.Println(query[0])
		url = fmt.Sprintf("https://api.tvmaze.com/singlesearch/shows?q=%s", query[1])
	} else if query[0] == "id" {
		fmt.Println(query[0])
		url = fmt.Sprintf("https://api.tvmaze.com/shows/%s", query[1])
	} else {
		fmt.Println("Unexpected URL query!")
		return
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &movies)
	json.NewEncoder(w).Encode(movies)
}