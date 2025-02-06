package scraper

import (
	"fmt"
	"os"
	"testing"
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

	text, err := ExtractText(htmlContent)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", text)
}
