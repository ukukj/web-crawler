package utils

import (
	"encoding/json"
	"fmt"
)

// LogJSON - 데이터를 prettify JSON  출력
func LogJSON(label string, data any) error {
	fmt.Printf("\n=== %s ===\n", label)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}

// LogJSONs - 슬라이스 -> LogJSON
func LogJSONs[T any](label string, data []T) error {
	for i, item := range data {
		if err := LogJSON(fmt.Sprintf("%s [%d]", label, i), item); err != nil {
			return err
		}
	}
	return nil
}
