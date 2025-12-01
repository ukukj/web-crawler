package crawler

type CrawlerConfig struct {
	Region   string
	Keywords []string
	MaxPages int
	Timeout  int
}

func NewConfig() CrawlerConfig {
	return CrawlerConfig{
		Region:   "",
		Keywords: []string{},
		MaxPages: 5,
		Timeout:  30,
	}
}
