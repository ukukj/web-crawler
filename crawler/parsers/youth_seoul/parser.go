package youth_seoul

import (
	"strings"
	"time"

	"web-crawler/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const (
	Name = "청년몽땅정보통"
	URL  = "https://youth.seoul.go.kr/infoData/plcyInfo/view.do?plcyBizId=V202500054&tab=001&key=&sc_detailAt=&pageIndex=1&orderBy=regYmd+desc&blueWorksYn=N&tabKind=001&sw="
	// URL  = "https://youth.seoul.go.kr"
)

type WelfareInfo struct {
	Name         string
	SupervisedBy string
	Description  string
	Contents     string
	Category     models.Category
	Span         *string
	Scale        *string
}

func SetupParser(c *colly.Collector, results *[]models.WelfareModel) {
	c.OnHTML(".policy-detail", func(elem *colly.HTMLElement) {
		info := GetWelfareInfo(elem)

		welfare := models.WelfareModel{
			BaseModel:     GetBaseModel(elem, info.Name),
			WelfareName:   info.Name,
			SupervisedBy:  info.SupervisedBy,
			Description:   info.Description,
			Contents:      info.Contents,
			Category:      info.Category,
			Span:          info.Span,
			Scale:         info.Scale,
			Qualification: GetQualification(elem),
			Registration:  GetRegistration(elem),
		}
		*results = append(*results, welfare)
	})
}

func GetBaseModel(elem *colly.HTMLElement, title string) models.BaseModel {
	return models.BaseModel{
		Title:     title,
		Source:    Name,
		Ref:       elem.Request.URL.String(),
		CrawledAt: time.Now().Format(time.RFC3339),
	}
}

func GetWelfareInfo(elem *colly.HTMLElement) WelfareInfo {
	// caption에 "사업개요"가 포함된 테이블 찾기
	table := elem.DOM.Find(".form-table.form-resp-table").FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Find("caption").Text(), "사업개요")
	}).First().Find("tbody")

	categoryText := strings.TrimSpace(table.Find("tr:nth-child(1) td:nth-child(2)").Text())

	return WelfareInfo{
		Name:         strings.TrimSpace(elem.ChildText(".top .lf strong.title")),
		SupervisedBy: strings.TrimSpace(table.Find("tr:nth-child(1) td:nth-child(4)").Text()),
		Description:  strings.TrimSpace(table.Find("tr:nth-child(2) td").Text()),
		Contents:     strings.TrimSpace(table.Find("tr:nth-child(3) td").Text()),
		Category:     parseCategory(categoryText),
		Span:         StringPtr(strings.TrimSpace(table.Find("tr:nth-child(4) td:nth-child(2)").Text())),
		Scale:        StringPtr(strings.TrimSpace(table.Find("tr:nth-child(5) td:nth-child(2)").Text())),
	}
}

func parseCategory(text string) models.Category {
	// 마침표를 쉼표로 정규화
	text = strings.ReplaceAll(text, ".", ",")

	switch text {
	case "일자리":
		return models.CategoryJob
	case "주거":
		return models.CategoryHousing
	case "교육":
		return models.CategoryEducation
	case "복지,문화":
		return models.CategoryWelfare
	case "참여,권리":
		return models.CategoryParticipate
	default:
		return models.CategoryJob
	}
}

func GetQualification(elem *colly.HTMLElement) models.Qualification {
	// caption에 "신청자격"이 포함된 테이블 찾기
	table := elem.DOM.Find(".form-table.form-resp-table").FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Find("caption").Text(), "신청자격")
	}).First().Find("tbody")

	ageText := strings.TrimSpace(table.Find("tr:nth-child(1) td:nth-child(2)").Text())
	education := strings.TrimSpace(table.Find("tr:nth-child(2) td:nth-child(2)").Text())
	major := strings.TrimSpace(table.Find("tr:nth-child(2) td:nth-child(4)").Text())
	employmentText := strings.TrimSpace(table.Find("tr:nth-child(3) td:nth-child(2)").Text())
	restriction := strings.TrimSpace(table.Find("tr:nth-child(5) td").Text())

	return models.Qualification{
		AgeRange:    ParseAgeRange(ageText),
		Education:   education,
		Major:       major,
		Employment:  ParseEmployment(employmentText),
		Restriction: restriction,
	}
}

func GetRegistration(elem *colly.HTMLElement) *models.Registration {
	// caption에 "신청방법"이 포함된 테이블 찾기
	table := elem.DOM.Find(".form-table.form-resp-table").FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Find("caption").Text(), "신청방법")
	}).First().Find("tbody")

	process := strings.TrimSpace(table.Find("tr:nth-child(1) td").Text())
	announcement := strings.TrimSpace(table.Find("tr:nth-child(2) td").Text())
	documents := strings.TrimSpace(table.Find("tr:nth-child(3) td").Text())
	url := strings.TrimSpace(table.Find("tr:nth-child(4) td a").AttrOr("href", ""))

	if process == "" && announcement == "" && documents == "" && url == "" {
		return nil
	}

	return &models.Registration{
		Process:      process,
		Announcement: announcement,
		Documents:    documents,
		URL:          url,
	}
}

func ParseAgeRange(text string) models.AgeRange {
	if strings.Contains(text, "19") && strings.Contains(text, "24") {
		return models.Age19to24
	}
	if strings.Contains(text, "25") && strings.Contains(text, "29") {
		return models.Age25to29
	}
	if strings.Contains(text, "30") && strings.Contains(text, "34") {
		return models.Age30to34
	}
	if strings.Contains(text, "35") && strings.Contains(text, "39") {
		return models.Age35to39
	}
	return models.Age19to24
}

func ParseEmployment(text string) models.EmploymentStatus {
	switch text {
	case "재직자":
		return models.EmploymentEmployed
	case "자영업자":
		return models.EmploymentSelfEmployed
	case "미취업자":
		return models.EmploymentUnemployed
	case "프리랜서":
		return models.EmploymentFreelancer
	case "일용근로자":
		return models.EmploymentDailyWorker
	case "(예비)창업자":
		return models.EmploymentPreStartup
	case "단기근로자":
		return models.EmploymentShortTerm
	case "영농종사자":
		return models.EmploymentFarmer
	default:
		return models.EmploymentAny
	}
}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
