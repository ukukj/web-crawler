package main

import (
	"fmt"

	"web-crawler/consts"
)

func main() {
	fmt.Println("=== 채용 사이트 ===")
	for _, site := range consts.Recruits {
		fmt.Printf("이름: %s, URL: %s\n", site.Name, site.URL)
	}

	fmt.Println("\n=== 복지 사이트 ===")
	for _, site := range consts.Welfare {
		fmt.Printf("이름: %s, URL: %s\n", site.Name, site.URL)
	}
}
