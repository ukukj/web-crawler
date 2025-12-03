package jobkorea

import (
	"time"

	"web-crawler/models"

	"github.com/gocolly/colly/v2"
)

const (
	Name = "잡코리아"
	URL  = "https://www.jobkorea.co.kr"
)

// SetupParser registers HTML parsing callbacks for JobKorea
func SetupParser(c *colly.Collector, results *[]models.JobModel) {
	// TODO: 실제 잡코리아 HTML 구조에 맞게 selector 수정 필요
	c.OnHTML(".job-list-item", func(elem *colly.HTMLElement) {
		job := ParseJobItem(elem)
		*results = append(*results, job)
	})
}

// GetTitle extracts job title from HTML element
func GetTitle(elem *colly.HTMLElement) string {
	// TODO: 실제 selector로 변경
	return elem.ChildText(".job-title")
}

// GetCompany extracts company name from HTML element
func GetCompany(elem *colly.HTMLElement) string {
	// TODO: 실제 selector로 변경
	return elem.ChildText(".company-name")
}

// GetJobURL extracts job posting URL from HTML element
func GetJobURL(elem *colly.HTMLElement) string {
	// TODO: 실제 속성명으로 변경
	return elem.Attr("href")
}

// GetLocation extracts job location from HTML element
func GetLocation(elem *colly.HTMLElement) string {
	// TODO: 실제 selector로 변경
	return elem.ChildText(".job-location")
}

// ParseJobItem assembles extracted data into JobModel
func ParseJobItem(elem *colly.HTMLElement) models.JobModel {
	return models.JobModel{
		BaseModel: models.BaseModel{
			Title:     GetTitle(elem),
			Source:    Name,
			Ref:       GetJobURL(elem),
			CrawledAt: time.Now().Format(time.RFC3339),
		},
		// TODO: JobModel에 필드 추가되면 여기도 추가
		// Company:  GetCompany(elem),
		// Location: GetLocation(elem),
	}
}
