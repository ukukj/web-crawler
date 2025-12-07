package repository

import (
	"database/sql"
	"time"

	"web-crawler/models"
)

func InsertJobs(db *sql.DB, jobs []models.JobModel) error {
	const upsertQuery = `
	INSERT INTO jobs (
		title, source, ref,
		company_name, work_place, career, education,
		crawled_at, created_at, is_active
	) VALUES (
		?, ?, ?,
		?, ?, ?, ?,
		?, ?, ?
	)
	ON DUPLICATE KEY UPDATE
		company_name = VALUES(company_name),
		work_place   = VALUES(work_place),
		career       = VALUES(career),
		education    = VALUES(education),
		crawled_at   = VALUES(crawled_at),
		is_active    = VALUES(is_active);
	`

	now := time.Now()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 1. 모든 공고 비활성화 (이번 크롤링 기준)
	if _, err := tx.Exec(`UPDATE jobs SET is_active = 0`); err != nil {
		_ = tx.Rollback()
		return err
	}

	// 2. 이번에 크롤링해 온 공고들을 upsert + is_active = 1
	stmt, err := tx.Prepare(upsertQuery)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, j := range jobs {
		_, err := stmt.Exec(
			j.Title,
			j.Source,
			j.Ref,
			j.CompanyName,
			j.WorkPlace,
			j.Career,
			j.Education,
			now,  // crawled_at
			now,  // created_at (신규 insert일 때만 의미 있음)
			true, // is_active
		)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	// 3. 아직도 is_active = 0 인 애들 = 이번 크롤링에 안 나온 공고 → 삭제
	if _, err := tx.Exec(`DELETE FROM jobs WHERE is_active = 0`); err != nil {
		_ = tx.Rollback()
		return err
	}

	// 4. 커밋
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
