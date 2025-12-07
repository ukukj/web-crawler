package saramin

import (
	"strings"
	"time"

	"web-crawler/models"

	"github.com/gocolly/colly/v2"
)

const (
	Name = "사람인"
	URL  = "https://www.saramin.co.kr/zf_user/jobs/list/domestic?loc_cd=102060&cat_kewd=84&panel_type=&search_optional_item=n&search_done=y&panel_count=y&preview=y"
)

// 사람인 전용 parser
func SetupParser(c *colly.Collector, results *[]models.JobModel) {
	c.OnHTML("div.list_body div.list_item", func(elem *colly.HTMLElement) {
		job := Parse(elem)
		if strings.TrimSpace(job.BaseModel.Title) == "" {
			return
		}
		*results = append(*results, job)
	})
}

// Parse : 사람인 채용 공고 하나 → JobModel 로 변환
func Parse(elem *colly.HTMLElement) models.JobModel {
	// 제목: div.notification_info > div.job_tit > a.str_tit
	title := elem.ChildAttr(".notification_info .job_tit a.str_tit", "title")
	if title == "" {
		title = elem.ChildText(".notification_info .job_tit a.str_tit")
	}
	title = strings.TrimSpace(title)

	// 상세 링크
	href := elem.ChildAttr(".notification_info .job_tit a.str_tit", "href")
	href = strings.TrimSpace(href)

	fullURL := ""
	if href != "" {
		if strings.HasPrefix(href, "http") {
			fullURL = href
		} else {
			fullURL = "https://www.saramin.co.kr" + href
		}
	}

	company := strings.TrimSpace(elem.ChildText(".company_nm a.str_tit"))
	workPlace := strings.TrimSpace(elem.ChildText(".recruit_info .work_place"))
	career := strings.TrimSpace(elem.ChildText(".recruit_info .career"))
	education := strings.TrimSpace(elem.ChildText(".recruit_info .education"))

	return models.JobModel{
		BaseModel: models.BaseModel{
			Title:     title,
			Source:    Name,
			Ref:       fullURL,
			CrawledAt: time.Now().Format(time.RFC3339),
		},
		CompanyName: company,
		WorkPlace:   workPlace,
		Career:      career,
		Education:   education,
	}
}
