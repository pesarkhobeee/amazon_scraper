// Package docscraper scrapes the page using doc attributes
package docscraper

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/pesarkhobeee/amazon_scraper/internal/model"
	"github.com/pesarkhobeee/amazon_scraper/internal/scraper"
)

var _ scraper.MovieParser = &Scraper{}

type Scraper struct{}

func (s Scraper) Parse(ctx context.Context, content string) (*model.MovieInformation, error) {
	if strings.Contains(content, "To discuss automated access to Amazon data please contact") {
		return nil, errors.New("blocked by Amazon")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	var movieInformation model.MovieInformation
	movieInformation.Title = doc.Find("h1[data-automation-id='title']").Text()

	movieInformation.ReleaseYear, err = strconv.Atoi(doc.Find("span[data-automation-id='release-year-badge']").Text())
	if err != nil {
		return nil, err
	}

	var actors_bare_html = doc.Find("#btf-product-details > div.\\+AZpnL > dl:nth-child(6) > dd")
	actors_bare_html.Find("a").Each(func(i int, s *goquery.Selection) {
		movieInformation.Actors = append(movieInformation.Actors, s.Text())
	})

	doc.Find("div[data-testid='packshot'] > a").Each(func(i int, s *goquery.Selection) {
		id := s.AttrOr("href", "")
		if substr := strings.Split(id, "/"); len(substr) >= 5 {
			id = substr[4]
		}
		movieInformation.SimilarIds = append(movieInformation.SimilarIds, id)
	})

	movieInformation.Poster = doc.Find("img[data-testid='base-image']").AttrOr("src", "")

	return &movieInformation, nil
}
