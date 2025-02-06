package httpfetcher

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// List of real User-Agents
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:112.0) Gecko/20100101 Firefox/112.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
}

// Function to get a random User-Agent
func getRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness
	return userAgents[rand.Intn(len(userAgents))]
}

func NewRequestWIthUserAgent(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// Set randomized User-Agent
	req.Header.Set("User-Agent", getRandomUserAgent())
	req.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, image/webp")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	return req, nil
}
