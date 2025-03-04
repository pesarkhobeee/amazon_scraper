package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pesarkhobeee/amazon_scraper/internal/handler"
	"github.com/pesarkhobeee/amazon_scraper/internal/scraper"
	"github.com/pesarkhobeee/amazon_scraper/internal/server"
)

/*
* Check this url
* http://localhost:8080/movie/amazon/B00KY1U7GM
 */

// Generic TODO
// 1. using interface to get the data from the scraper
// 2. using context to timeout the request
// 3. using goroutine to scrape the data
// 4. using channel to send the data to the server
// 5. using gorilla mux to handle the request
// 6. using goquery to parse the html
// 7. using gorilla mux to handle the request
// 8. using goquery to parse the html
// 9. using gorilla mux to handle the request
// 10. using goquery to parse the html

// TODO:
// receive multiple amazon ids (e.g. in the query string)
// use the func but with a wrapper that runs them in separate goroutines and use a channel to send the data to the server
// if you want you can limit number of workers
// communicate between goroutines with channel
// channel, sync.WaitGroup, sync.Mutex
// https://pkg.go.dev/golang.org/x/sync/errgroup

func main() {
	// log_level := os.Getenv("LOG_LEVEL")
	// if log_level == "" {
	// 	log_level = "info"
	// }

	// 2. Set the log level

	log.Printf("Starting the server on port %d...", port)

	// 1. Run the server
	router, err := newRouter(nil, http.DefaultClient)
	if err != nil {
		panic(err)
	}
	srv := server.NewServer(port, router)
	log.Println(srv.ListenAndServe())
}

func newRouter(scraper scraper.MovieParser, httpClient *http.Client) (http.Handler, error) {
	handler, err := handler.NewMovieScraper(baseAddress, scraper, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new handler: %w", err)
	}

	router := mux.NewRouter()
	router.
		HandleFunc("/movie/amazon/{amazon_id}", handler.GetAmazonMovieInformation).
		Methods(http.MethodGet)
	return router, nil
}

const (
	port        = 8080
	baseAddress = "https://www.amazon.de"
)
