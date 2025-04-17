package utilities

import "time"

var startTime time.Time

func StartTiming() {
	startTime = time.Now()
}

func CalculateResponseTime() int {
	return int(time.Since(startTime).Milliseconds())
}
