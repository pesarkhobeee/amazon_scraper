package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ScrapeAmazonMovieInformation(amazon_id string) string {
	url := fmt.Sprintf("https://www.amazon.de/gp/product/%s", amazon_id)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Set the User-Agent header to mimic Firefox
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:112.0) Gecko/20100101 Firefox/112.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
