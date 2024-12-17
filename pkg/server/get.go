package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pesarkhobeee/amazon_scraper/pkg/scraper"
)

// RunServer runs the server
func RunServer() {

	router := mux.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.HandleFunc("/movie/amazon/{amazon_id}", getAmazonMovieInformation).Methods("GET")

	log.Fatal(srv.ListenAndServe())
}

func getAmazonMovieInformation(w http.ResponseWriter, r *http.Request) {
	amazon_id := mux.Vars(r)["amazon_id"]
	content := scraper.ScrapeAmazonMovieInformation(amazon_id)
	w.Write([]byte(content))
}
