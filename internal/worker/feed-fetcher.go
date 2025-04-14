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

		defer func() {
			if r := recover(); r != nil {
				log.Printf("[FeedFetcher] Recovered from panic: %v", r)
			}
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[FeedFetcher] Error fetching url %s: %s", url, err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[FeedFetcher] Error when reading response body: %v", err)
			return
		}

		var parsed model.FetchResponse
		if err = json.Unmarshal(body, &parsed); err != nil {
			log.Printf("[FeedFetcher] Error when parsing JSON: %v", err)
			return
		}

		if parsed.Status != "success" {
			log.Printf("[FeedFetcher] Error fetching url %s: %s", url, parsed.Status)
			return
		}

		err = f.productRepository.SaveProducts(ctx, parsed.Data.MachineProducts)
		if err != nil {
			log.Printf("[FeedFetcher] Error when saving products: %v", err)
			return
		}

		log.Printf("[FeedFetcher] Successfully fetched and saved %d products", len(parsed.Data.MachineProducts))
	}
}
