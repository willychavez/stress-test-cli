package domain

import "time"

type Report struct {
	TotalRequests      int
	SuccessRequests    int
	StatusDistribution map[int]int
	TotalTime          time.Duration
}
