package db

// 이 패키지 import 해서 사용
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New() *sql.DB {
	// root: crawler: crawlerpw, DB: crawler
	dsn := fmt.Sprintf(
		"crawler:crawler1212!@#@tcp(127.0.0.1:13306)/crawler?parseTime=true&charset=utf8mb4&loc=Asia%%2FSeoul",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}

	// 커넥션 옵션
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping DB:", err)
	}

	return db
}
