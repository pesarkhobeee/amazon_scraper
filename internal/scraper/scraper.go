package scraper

import (
	"context"

	"github.com/pesarkhobeee/amazon_scraper/internal/model"
)

type MovieParser interface {
	Parse(ctx context.Context, content string) (*model.MovieInformation, error)
}
