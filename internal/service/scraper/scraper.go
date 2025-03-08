package scraper

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pesarkhobeee/amazon_scraper/internal/model"
	"github.com/pesarkhobeee/amazon_scraper/internal/parser"
	"github.com/pesarkhobeee/amazon_scraper/pkg/httpfetcher"
)

type Service struct {
	parser      parser.MovieParser
	baseAddress *url.URL
	client      *http.Client
}

func NewService(
	address string,
	parser parser.MovieParser,
	client *http.Client,
) (*Service, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}
	return &Service{
		parser:      parser,
		baseAddress: u,
		client:      client,
	}, nil
}

func (s *Service) Scrape(ctx context.Context, amazonID string) (*model.MovieInformation, error) {
	reqURL := s.baseAddress.JoinPath("gp", "product", amazonID).String()
	req, err := httpfetcher.NewRequestWIthUserAgent(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client error: %w", err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	info, err := s.parser.Parse(ctx, string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to scrape the content: %w", err)
	}

	return info, nil
}
