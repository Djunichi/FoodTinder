package worker

import (
	"context"
	"encoding/json"
	"food-tinder/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

type ProductSaver interface {
	SaveProducts(ctx context.Context, products []model.MachineProduct) error
}

type FeedFetcher struct {
	productRepository ProductSaver
}

func NewFeedFetcher(repo ProductSaver) *FeedFetcher {
	return &FeedFetcher{productRepository: repo}
}

func (f *FeedFetcher) FetchFeed(url string) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("Error fetching url %s: %s", url, err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error when reading responce body: %v", err)
		}

		var parsed model.FetchResponse
		if err := json.Unmarshal(body, &parsed); err != nil {
			log.Fatalf("Error when parsing JSON: %v", err)
		}

		err = f.productRepository.SaveProducts(ctx, parsed.Data.MachineProducts)
		if err != nil {
			log.Fatalf("Error when saving products: %v", err)
		}
	}
}
