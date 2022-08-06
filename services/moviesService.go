package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
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

var movieByTitle interface{}

func SearchMovies(w http.ResponseWriter, r *http.Request) {
	var url string

	query := strings.Split(r.URL.RawQuery, "=")

	if query[0] == "movie" {
		url = fmt.Sprintf("https://api.tvmaze.com/search/shows?q=%s", query[1])
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
		return
	} else if query[0] == "title" {
		url = fmt.Sprintf("https://api.tvmaze.com/singlesearch/shows?q=%s", query[1])

		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &movieByTitle)
		json.NewEncoder(w).Encode(movieByTitle)
		return
	} else {
		fmt.Println("Unexpected URL query!")
		return
	}
}

var moviesById interface{}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	url := fmt.Sprintf("https://api.tvmaze.com/shows/%s", id)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &moviesById)
	json.NewEncoder(w).Encode(moviesById)
}
