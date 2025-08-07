package jobs

import (
	"fmt"
	"time"
)

func Worker(id int) {
	for {
		fmt.Printf("INFO: Worker %d starting data collection\n", id)
		start := time.Now()

		Data()

		duration := time.Since(start)
		fmt.Printf("INFO: Worker %d completed in %v\n", id, duration)
	}
}
