package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/pesarkhobeee/amazon_scraper/internal/scraper"
	"github.com/pesarkhobeee/amazon_scraper/pkg/httpfetcher"
)

type MovieScraper struct {
	parser      scraper.MovieParser
	baseAddress *url.URL
	client      *http.Client
}

func NewMovieScraper(
	address string,
	parser scraper.MovieParser,
	client *http.Client,
) (*MovieScraper, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}
	return &MovieScraper{
		parser:      parser,
		baseAddress: u,
		client:      client,
	}, nil
}

func (h *MovieScraper) GetAmazonMovieInformation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	amazonID := mux.Vars(r)["amazon_id"]

	info, err := h.getAmazonMovieInformation(ctx, amazonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(info)
}

func (h *MovieScraper) getAmazonMovieInformation(ctx context.Context, amazonID string) (*scraper.MovieInformation, error) {
	reqURL := h.baseAddress.JoinPath("gp", "product", amazonID).String()
	req, err := httpfetcher.NewRequestWIthUserAgent(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client error: %w", err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	info, err := h.parser.Parse(ctx, string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to scrape the content: %w", err)
	}

	return info, nil
}
