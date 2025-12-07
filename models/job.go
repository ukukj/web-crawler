package models

type JobModel struct {
	BaseModel

	CompanyName string // 회사명
	WorkPlace   string // 근무지
	Career      string // 경력/고용형태
	Education   string // 학력
}
