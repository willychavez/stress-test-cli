package main

import (
	"log"
	"time"

	"github.com/willychavez/stress-test-cli/internal/config"
	"github.com/willychavez/stress-test-cli/internal/httpclient"
	"github.com/willychavez/stress-test-cli/internal/presentation"
	"github.com/willychavez/stress-test-cli/internal/usecase"
)

func main() {
	// Parse CLI flags
	cliConfig, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	// Create HTTP client
	httpClient := httpclient.NewHttpClient()

	// Start time tracking
	start := time.Now()

	log.Printf("Starting stress test to %s with %d requests and %d concurrency\n", cliConfig.URL, cliConfig.TotalRequests, cliConfig.Concurrency)

	// Execute requests
	results := usecase.ExecuteRequests(httpClient, cliConfig.URL, cliConfig.TotalRequests, cliConfig.Concurrency)

	// Generate report
	report := usecase.GenerateReport(results, start)

	// Print report
	presentation.PrintReport(report)
}
