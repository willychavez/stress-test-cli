package presentation

import (
	"fmt"

	"github.com/willychavez/stress-test-cli/internal/domain"
)

func PrintReport(report domain.Report) {
	fmt.Println("\n================= Report =================")
	fmt.Printf("Total execution time: %s\n", report.TotalTime)
	fmt.Printf("Total requests made: %d\n", report.TotalRequests)
	fmt.Printf("Total requests with status 200: %d\n", report.SuccessRequests)
	fmt.Println("Status code distribution:")
	for status, count := range report.StatusDistribution {
		fmt.Printf("  %d: %d\n", status, count)
	}
	fmt.Println("============================================")
}
