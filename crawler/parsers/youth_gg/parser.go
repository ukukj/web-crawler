package youth_gg

import (
	"time"

	"web-crawler/models"

	"github.com/gocolly/colly/v2"
)

const (
	Name = "경기청년포탈"
	URL  = "https://youth.gg.go.kr"
)

// Parse parses welfare information from Gyeonggi Youth Portal
func Parse(elem *colly.HTMLElement) models.WelfareModel {
	// TODO: HTML 구조 분석 후 작성
	return models.WelfareModel{
		BaseModel: models.BaseModel{
			Title:     "", // TODO
			Source:    Name,
			Ref:       "", // TODO
			CrawledAt: time.Now().Format(time.RFC3339),
		},
	}
}
