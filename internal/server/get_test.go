package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pesarkhobeee/amazon_scraper/internal/scraper"
)

func TestGetAmazonMovieInformation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	amazon_id := "B00K19SD8Q"
	htmlContent, err := scraper.ScrapeAmazonMovieInformation(ctx, amazon_id)
	if err != nil {
		t.Error(err)
	}

	if htmlContent == "" {
		t.Error("Could not get the content of the page")
	}

	text, err := scraper.ExtractText(htmlContent)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(text)
}
