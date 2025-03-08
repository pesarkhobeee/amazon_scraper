package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pesarkhobeee/amazon_scraper/internal/model"
)

type MovieScraperService interface {
	Scrape(ctx context.Context, amazonID string) (*model.MovieInformation, error)
}

type MovieScraper struct {
	service MovieScraperService
}

func NewMovieScraper(service MovieScraperService) *MovieScraper {
	return &MovieScraper{
		service: service,
	}
}

func (h *MovieScraper) GetAmazonMovieInformation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	amazonID := mux.Vars(r)["amazon_id"]

	info, err := h.service.Scrape(ctx, amazonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(info)
}
