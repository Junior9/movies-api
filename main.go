package main

import (
	"log"
	"net/http"

	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/handler"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
)

func main() {
	movieSearcher := &movies.APIMovieSearcher{
		APIKey: "47494e7e",
		URL:    "https://www.omdbapi.com/",
	}

	handler := handler.NewHandler(movieSearcher)
	//"Serve Starting..."
	log.Fatal(http.ListenAndServe(":3000", handler))
}
