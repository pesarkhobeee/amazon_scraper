package server

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	amazon_id := mux.Vars(r)["amazon_id"]
	fmt.Println(amazon_id)
	htmlContent := scraper.ScrapeAmazonMovieInformation(ctx, amazon_id)

	if htmlContent == "" {
		log.Fatal("Could not get the content of the page")
	}

	text, err := scraper.ExtractText(htmlContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
