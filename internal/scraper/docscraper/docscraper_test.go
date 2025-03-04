package docscraper_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/pesarkhobeee/amazon_scraper/internal/model"
	"github.com/pesarkhobeee/amazon_scraper/internal/scraper/docscraper"
)

func openFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}

func TestGetAmazonMovieInformation(t *testing.T) {
	htmlContent := openFile("test_data/amazon_movie_information.html")

	if htmlContent == "" {
		t.Error("Could not get the content of the page")
	}

	scp := &docscraper.Scraper{}
	got, err := scp.Parse(context.TODO(), htmlContent)
	if err != nil {
		t.Error(err)
	}

	want := &model.MovieInformation{
		Title:       "Um Jeden Preis [dt./OV]",
		ReleaseYear: 2012,
		Actors:      []string{"Dennis Quaid", "Zac Efron", "Kim Dickens", "Heather Graham"},
		Poster:      "./Watch Um Jeden Preis [dt._OV] _ Prime Video_files/b86faa650d7a4f77afb5b01577ec18dab930b4b2b0d2ccb484145599ebe7e896._SX1080_FMpng_.png",
		SimilarIds:  []string{"en", "en", "en", "en", "en", "en", "en", "en", "en", "en", "en"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}
