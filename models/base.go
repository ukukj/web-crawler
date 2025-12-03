package models

type BaseModel struct {
	ID        string
	Title     string
	Source    string
	Ref       string
	CrawledAt string // RFC3339
	CreatedAt string // RFC3339
	IsActive  bool
}
