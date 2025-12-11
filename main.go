package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"web-crawler/crawler"
	"web-crawler/crawler/parsers/saramin"
	"web-crawler/crawler/parsers/youth_seoul"
	"web-crawler/db"
	"web-crawler/models"
	"web-crawler/repository"
	"web-crawler/utils"
)

func main() {
	fmt.Println("Starting Crawling Jobs")

	// DB 연결 ------------------------------------
	database := db.New()
	defer database.Close()
	// ---------------------------------------------

	var wg sync.WaitGroup
	cfg := crawler.NewConfig()

	JobResults := make(chan []models.JobModel, 1)
	WelfareResults := make(chan []models.WelfareModel, 2)

	// 크롤링할 "사이트 개수"만큼 Add
	wg.Add(2)

	// 청년몽땅 정보포털
	go func() {
		defer wg.Done()
		fmt.Println("Begin Crawling:", youth_seoul.URL)
		result, err := crawler.Crawl(youth_seoul.Name, youth_seoul.URL, cfg, youth_seoul.SetupParser)
		if err != nil {
			log.Fatal("Crawl Failed", youth_seoul.Name, err)
		}
		WelfareResults <- result
	}()

	// 사람인
	go func() {
		defer wg.Done()
		fmt.Println("Begin Crawling:", saramin.URL)
		result, err := crawler.Crawl(saramin.Name, saramin.URL, cfg, saramin.SetupParser)
		if err != nil {
			log.Fatal("Crawl Failed", saramin.Name, err)
		}
		JobResults <- result
	}()

	// 채널 닫기 goroutine
	go func() {
		wg.Wait()
		close(JobResults)
		close(WelfareResults)
	}()

	allJobs := utils.CollectResults(JobResults)
	allWelfares := utils.CollectResults(WelfareResults)

	fmt.Printf("크롤링 완료. Job: %d개, Welfare: %d개\n", len(allJobs), len(allWelfares))

	// 파싱 결과 JSON 출력 (첫 번째만)
	if len(allJobs) > 0 {
		if err := utils.LogJSON("Job 파싱 결과", allJobs[0]); err != nil {
			log.Println("JSON 출력 실패:", err)
		}
	}

	if len(allWelfares) > 0 {
		if err := utils.LogJSON("Welfare 파싱 결과", allWelfares[0]); err != nil {
			log.Println("JSON 출력 실패:", err)
		}
	}

	// DB 저장 ------------------------------------
	if len(allJobs) > 0 {
		if err := repository.InsertJobs(database, allJobs); err != nil {
			log.Fatal("InsertJobs error:", err)
		}
		fmt.Println("DB 저장 완료:", len(allJobs), "건")
	} else {
		fmt.Println("저장할 Job 데이터가 없습니다.")
	}
	// ---------------------------------------------
}
