package scraper

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pesarkhobeee/amazon_scraper/pkg/httpfetcher"
)

func ScrapeAmazonMovieInformation(ctx context.Context, amazonID string) (string, error) {
	url := fmt.Sprintf("https://www.amazon.de/gp/product/%s", amazonID)

	client := &http.Client{}
	req, err := httpfetcher.NewRequestWIthUserAgent(ctx, "GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
