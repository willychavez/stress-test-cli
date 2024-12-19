package usecase

import (
	"net/http"
	"sync"
	"time"

	"github.com/willychavez/stress-test-cli/internal/domain"
)

func ExecuteRequests(client *http.Client, url string, totalRequests, concurrency int) chan domain.Result {
	results := make(chan domain.Result, totalRequests)
	var wg sync.WaitGroup
	var mu sync.Mutex
	requestCount := 0

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if requestCount >= totalRequests {
					mu.Unlock()
					return
				}
				requestCount++
				mu.Unlock()

				resp, err := client.Get(url)

				if err != nil {
					results <- domain.Result{Err: err}
					continue
				}
				results <- domain.Result{StatusCode: resp.StatusCode}
				resp.Body.Close()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func GenerateReport(results chan domain.Result, start time.Time) domain.Report {
	totalRequests := 0
	successRequests := 0
	statusDistribution := make(map[int]int)

	for result := range results {
		totalRequests++
		if result.Err != nil {
			continue
		}
		if result.StatusCode == http.StatusOK {
			successRequests++
		}
		statusDistribution[result.StatusCode]++
	}

	totalTime := time.Since(start)

	return domain.Report{
		TotalRequests:      totalRequests,
		SuccessRequests:    successRequests,
		StatusDistribution: statusDistribution,
		TotalTime:          totalTime,
	}
}
