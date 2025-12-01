package models

import "time"

type BaseModel struct {
	ID        string
	Title     string
	Source    string
	Ref       string
	CrawledAt time.Time
	CreatedAt time.Time
	IsActive  bool
}
