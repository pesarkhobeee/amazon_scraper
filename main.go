package main

import (
	"log"

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

	log.Println("Starting the server...")

	// 1. Run the server
	server.RunServer()
}
