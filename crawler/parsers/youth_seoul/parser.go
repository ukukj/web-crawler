package youth_seoul

import (
	"time"

	"web-crawler/models"

	"github.com/gocolly/colly/v2"
)

const (
	Name = "청년몽땅정보통"
	URL  = "https://youth.seoul.go.kr"
)

func SetupParser(c *colly.Collector, results *[]models.WelfareModel) {
	// TODO: update selector based on actual HTML structure
	c.OnHTML(".welfare-list-item", func(elem *colly.HTMLElement) {
		welfare := ParseWelfareItem(elem)
		*results = append(*results, welfare)
	})
}

func GetTitle(elem *colly.HTMLElement) string {
	return elem.ChildText(".welfare-title")
}

func GetWelfareURL(elem *colly.HTMLElement) string {
	return elem.Attr("href")
}

func GetEligibility(elem *colly.HTMLElement) string {
	return elem.ChildText(".eligibility")
}

func GetDeadline(elem *colly.HTMLElement) string {
	return elem.ChildText(".deadline")
}

func ParseWelfareItem(elem *colly.HTMLElement) models.WelfareModel {
	return models.WelfareModel{
		BaseModel: models.BaseModel{
			Title:     GetTitle(elem),
			Source:    Name,
			Ref:       GetWelfareURL(elem),
			CrawledAt: time.Now().Format(time.RFC3339),
		},
	}
}
