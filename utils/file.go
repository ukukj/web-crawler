package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// SaveResultAsFile - Crawling 결과 Dump html로 저장
func SaveResultAsFile(filename string, content string) error {
	now := time.Now()

	// 디렉토리: 날짜만
	dateDir := now.Format("2006-01-02")
	targetDir := filepath.Join("dumps", dateDir)

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	// 파일명: 시간 추가
	timeStamp := now.Format("15-04-05")
	ext := filepath.Ext(filename)
	nameWithoutExt := filename[:len(filename)-len(ext)]
	fullFilename := fmt.Sprintf("%s_%s%s", nameWithoutExt, timeStamp, ext)

	fullPath := filepath.Join(targetDir, fullFilename)

	if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
		return err
	}

	fmt.Printf("File saved → %s\n", fullPath)
	return nil
}
