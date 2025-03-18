package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"

	handlermocks "github.com/pesarkhobeee/amazon_scraper/internal/handler/mocks"
	"github.com/pesarkhobeee/amazon_scraper/internal/model"
)

func TestMovieScraper_GetMultipleMovieInformation(t *testing.T) {
	serviceMock := handlermocks.NewMockMovieScraperService(t)
	serviceMock.On("Scrape", mock.Anything, "foo").Return(&model.MovieInformation{
		Title:       "foo",
		ReleaseYear: 1990,
	}, nil)
	serviceMock.On("Scrape", mock.Anything, "bar").Return(&model.MovieInformation{
		Title:       "bar",
		ReleaseYear: 1991,
	}, nil)
	h := NewMovieScraper(serviceMock)

	ts := httptest.NewServer(http.HandlerFunc(h.GetMultipleAmazonMovieInformation))
	defer ts.Close()

	url := fmt.Sprintf("%s?id=foo&id=bar", ts.URL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	want := []*model.MovieInformation{
		{Title: "foo", ReleaseYear: 1990},
		{Title: "bar", ReleaseYear: 1991},
	}
	var got []*model.MovieInformation
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

func TestMovieScraper_getMultipleAmazonMovieInformation(t *testing.T) {
	var unknownError = errors.New("unknown error")

	ids := []string{"foo", "bar", "fooz"}
	tt := map[string]struct {
		mockFunc func(*handlermocks.MockMovieScraperService)
		want     []*model.MovieInformation
		err      error
	}{
		"serviceError": {
			mockFunc: func(m *handlermocks.MockMovieScraperService) {
				m.On("Scrape", mock.Anything, "foo").Return(nil, unknownError)
				m.On("Scrape", mock.Anything, "bar").Return(&model.MovieInformation{}, nil)
				m.On("Scrape", mock.Anything, "fooz").Return(&model.MovieInformation{}, nil)
			},
			err: unknownError,
		},
		"happyCase": {
			mockFunc: func(m *handlermocks.MockMovieScraperService) {
				m.On("Scrape", mock.Anything, "foo").Return(&model.MovieInformation{
					Title:       "foo",
					ReleaseYear: 1990,
				}, nil)
				m.On("Scrape", mock.Anything, "bar").Return(&model.MovieInformation{
					Title:       "bar",
					ReleaseYear: 1991,
				}, nil)
				m.On("Scrape", mock.Anything, "fooz").Return(&model.MovieInformation{
					Title:       "fooz",
					ReleaseYear: 1992,
				}, nil)
			},
			want: []*model.MovieInformation{
				{Title: "foo", ReleaseYear: 1990},
				{Title: "bar", ReleaseYear: 1991},
				{Title: "fooz", ReleaseYear: 1992},
			},
		},
	}

	for testName, c := range tt {
		t.Run(testName, func(t *testing.T) {
			serviceMock := handlermocks.NewMockMovieScraperService(t)
			if f := c.mockFunc; f != nil {
				f(serviceMock)
			}
			h := NewMovieScraper(serviceMock)
			got, err := h.getAmazonMovieInfos(context.Background(), ids)

			if !errors.Is(err, c.err) {
				t.Errorf("unexpected error. want %s but got %s", c.err, err)
			}
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
