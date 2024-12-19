package config

import (
	"flag"
	"fmt"
)

type CLIConfig struct {
	URL           string
	TotalRequests int
	Concurrency   int
}

func ParseFlags() (CLIConfig, error) {
	url := flag.String("url", "", "URL of the service to be tested")
	totalRequests := flag.Int("requests", 100, "Total number of requests")
	concurrency := flag.Int("concurrency", 10, "Number of simultaneous calls")

	flag.Parse()

	if *url == "" {
		return CLIConfig{}, fmt.Errorf("URL is required")
	}

	return CLIConfig{
		URL:           *url,
		TotalRequests: *totalRequests,
		Concurrency:   *concurrency,
	}, nil
}
