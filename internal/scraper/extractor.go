package scraper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
*
* {

"title": "Um Jeden Preis",

"release_year": 2013,

"actors": ["Dennis Quaid","Zac Efron"],

"poster": "http://ecx.images-
amazon.com/images/I/51UZ8st2OdL._SX200_QL80_.jpg",

"similar_ids":
["B00SWDQPOC","B00RBPBO1G","B00S2EMECI","B00M5GH53M","B00IH8BA3S",
"B00M5JP1DA"]

}

*
*/

type MovieInformation struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

func ExtractText(htmlContent string) (*MovieInformation, error) {

	if strings.Contains(htmlContent, "To discuss automated access to Amazon data please contact") {
		return nil, errors.New("blocked by Amazon")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var movieInformation MovieInformation

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
