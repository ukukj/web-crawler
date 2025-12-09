package models

// 우선은 youth.seoul.go.kr 을 기준으로 모델 산정

// ENUMS

// Category - 정책 유형
type Category string

const (
	CategoryJob         Category = "일자리"
	CategoryHousing     Category = "주거"
	CategoryEducation   Category = "교육"
	CategoryWelfare     Category = "복지,문화"
	CategoryParticipate Category = "참여,권리"
)

// EmploymentStatus - 취업상태
type EmploymentStatus string

const (
	EmploymentEmployed     EmploymentStatus = "재직자"
	EmploymentSelfEmployed EmploymentStatus = "자영업자"
	EmploymentUnemployed   EmploymentStatus = "미취업자"
	EmploymentFreelancer   EmploymentStatus = "프리랜서"
	EmploymentDailyWorker  EmploymentStatus = "일용근로자"
	EmploymentPreStartup   EmploymentStatus = "(예비)창업자"
	EmploymentShortTerm    EmploymentStatus = "단기근로자"
	EmploymentFarmer       EmploymentStatus = "영농종사자"
	EmploymentAny          EmploymentStatus = "제한없음"
)

// AgeRange - 연령대
type AgeRange string

const (
	Age19to24 AgeRange = "19~24"
	Age25to29 AgeRange = "25~29"
	Age30to34 AgeRange = "30~34"
	Age35to39 AgeRange = "35~39"
)

// Models

type WelfareModel struct {
	BaseModel

	WelfareName   string        // 정책 이름
	SupervisedBy  string        // 주관 기관
	Description   string        // 정책 소개
	Contents      string        // 정책 내용
	Category      Category      // 정책 유형
	Span          *string       // 운영 기간 (optional)
	Scale         *string       // 지원규모 (optional)
	Qualification Qualification // 신청자격
	Registration  *Registration // 신청방법 (optional)
}

type Qualification struct {
	AgeRange    AgeRange         // 연령
	Education   string           // 학력
	Major       string           // 전공요건
	Employment  EmploymentStatus // 취업상태
	Restriction string           // 제한대상
}

type Registration struct {
	Process      string // 절차
	Announcement string // 심사 및 발표
	Documents    string // 제출 서류
	URL          string // 사이트
}
