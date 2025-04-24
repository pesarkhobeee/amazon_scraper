package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

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

func (h *MovieScraper) GetMultipleAmazonMovieInformation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		ids = r.URL.Query()["id"]
	)
	infos, err := h.getAmazonMovieInfos(ctx, ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(infos)
}

func (h *MovieScraper) getAmazonMovieInfos(ctx context.Context, ids []string) ([]*model.MovieInformation, error) {
	var (
		// make sure that we allocate the slice with same number of input
		// the purpose is we want to preserve the index of the input ids
		// if we only use append(..) then the order can be random since it's happening in separate goroutine
		res = make([]*model.MovieInformation, len(ids))

		// channel here is to communicate between goroutines
		// in this case we only need to pass the error
		// we close the channel once we already pass the last element to it
		// panic will happen if you pass another value to a closed channel
		// but it's fine if you receive from a closed channel
		ch = make(chan error)

		// wait group will block and wait until all of the wg.Add(x) is
		// responded by equal number of wg.Done()
		// here it's used to close the channel once we know that all necessary
		// values have been input into the channel
		wg sync.WaitGroup

		// global errors within the scope of this func
		// to be populated using errors.Join(...) func to capture all of the
		// occuring errors
		errs error
	)
	for idx, id := range ids {
		// run the scraper in separate goroutines and increment the waitgroup for each
		wg.Add(1)
		go h.scrapeMovieInformation(ctx, id, idx, ch, &wg, res)
	}
	go func() {
		// this will block in separate goroutine and close the channel once
		// all scrapers are done
		wg.Wait()
		close(ch)
	}()

	// this range will break the loop once channel ch is closed
	// otherwise it will keep listening to all errors passed into the channel
	for err := range ch {
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	if errs != nil {
		return nil, errs
	}
	return res, errs
}

func (h *MovieScraper) scrapeMovieInformation(
	ctx context.Context,
	id string,
	idx int,
	ch chan<- error,
	wg *sync.WaitGroup,
	res []*model.MovieInformation,
) {
	defer wg.Done()
	info, err := h.service.Scrape(ctx, id)
	if err != nil {
		ch <- fmt.Errorf("failed to scrape amazon id[%s]: %w", id, err)
		return
	}
	res[idx] = info
	ch <- nil
}
