package main

import (
	"fmt"
	"log"
	"sync"

	"web-crawler/crawler"
	"web-crawler/crawler/parsers/jobkorea"
	"web-crawler/crawler/parsers/youth_seoul"
	"web-crawler/models"
)

func main() {
	fmt.Println("Starting Crawling Jobs")

	var wg sync.WaitGroup
	cfg := crawler.NewConfig()

	JobResults := make(chan []models.JobModel, 2)
	WelfareResults := make(chan []models.WelfareModel, 2)

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Begin Crawling:", jobkorea.URL)
		result, err := crawler.Crawl(jobkorea.URL, cfg, jobkorea.SetupParser)
		if err != nil {
			log.Fatal("Crawl Failed", jobkorea.Name, err)
		}
		JobResults <- result
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Begin Crawling:", youth_seoul.URL)
		result, err := crawler.Crawl(youth_seoul.URL, cfg, youth_seoul.SetupParser)
		if err != nil {
			log.Fatal("Crawl Failed", youth_seoul.Name, err)
		}
		WelfareResults <- result
	}()

	go func() {
		wg.Wait()
		close(JobResults)
		close(WelfareResults)
	}()

	allJobs := []models.JobModel{}
	for jobs := range JobResults {
		allJobs = append(allJobs, jobs...)
	}

	fmt.Printf("크롤링 완료. 결과: %d개\n", len(allJobs))
}
