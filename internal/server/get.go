package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pesarkhobeee/amazon_scraper/internal/scraper"
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

	log.Println(srv.ListenAndServe())
}

func getAmazonMovieInformation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	amazon_id := mux.Vars(r)["amazon_id"]
	fmt.Println(amazon_id)
	htmlContent, err := scraper.ScrapeAmazonMovieInformation(ctx, amazon_id)
	if err != nil {
		http.Error(w, "Could not scrape the page: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if htmlContent == "" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	text, err := scraper.ExtractText(htmlContent)
	if err != nil {
		http.Error(w, "Could not extract the text from the page", http.StatusInternalServerError)
		return
	}
	fmt.Println(text)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(text)
}
