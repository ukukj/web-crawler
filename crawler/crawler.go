package crawler

import (
	"log"

	"web-crawler/utils"

	"github.com/gocolly/colly/v2"
)

// Crawl fetches and parses data from the given URL using the provided parser setup function
func Crawl[T any](siteName string, url string, cfg CrawlerConfig, setupParser func(*colly.Collector, *[]T)) ([]T, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; WebCrawler/1.0)"),
	)

	results := []T{}

	// HTML 보기 위해 파일로 저장
	c.OnResponse(func(r *colly.Response) {
		_ = utils.SaveResultAsFile(siteName+".html", string(r.Body))
	})

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
