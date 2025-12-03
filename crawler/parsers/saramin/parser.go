package saramin

import (
	"time"

	"web-crawler/models"

	"github.com/gocolly/colly/v2"
)

const (
	Name = "사람인"
	URL  = "https://www.saramin.co.kr"
)

// Parse parses job postings from Saramin
func Parse(elem *colly.HTMLElement) models.JobModel {
	// TODO: HTML 구조 분석 후 작성
	return models.JobModel{
		BaseModel: models.BaseModel{
			Title:     "", // TODO
			Source:    Name,
			Ref:       "", // TODO
			CrawledAt: time.Now().Format(time.RFC3339),
		},
	}
}
