package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

	if r.URL.Query().Get("movie") != "" {
		fmt.Println("movie = " + r.URL.Query().Get("movie"))
		url = fmt.Sprintf("https://api.tvmaze.com/search/shows?q=%s", r.URL.Query().Get("movie"))
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
	} else if r.URL.Query().Get("title") != "" {
		fmt.Println("title = " + r.URL.Query().Get("title"))
		url = fmt.Sprintf("https://api.tvmaze.com/singlesearch/shows?q=%s", r.URL.Query().Get("title"))

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
		json.NewEncoder(w).Encode("Unexpected URL query!")
		return
	}
}

var moviesById interface{}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

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
