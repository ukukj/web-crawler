package utils

// CollectResults - 채널에서 결과를 받아 슬라이스로 변환
func CollectResults[T any](ch chan []T) []T {
	results := []T{}
	for items := range ch {
		results = append(results, items...)
	}
	return results
}
