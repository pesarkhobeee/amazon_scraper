package main

import (
	"log"

	"github.com/pesarkhobeee/amazon_scraper/pkg/server"
)

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

/*
* http://localhost:8080/movie/amazon/B00KY1U7GM
* TODO:
* check this https://pkg.go.dev/golang.org/x/net/html
 */
