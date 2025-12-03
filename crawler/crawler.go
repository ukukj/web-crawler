package crawler

import (
	"log"

	"github.com/gocolly/colly/v2"
)

// Crawl fetches and parses data from the given URL using the provided parser setup function
func Crawl[T any](url string, cfg CrawlerConfig, setupParser func(*colly.Collector, *[]T)) ([]T, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; WebCrawler/1.0)"),
	)

	results := []T{}

	// 요청 전 로깅
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
	})

	// 에러 핸들링
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	setupParser(c, &results)

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return results, nil
}
