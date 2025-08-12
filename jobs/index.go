package jobs

import (
	"fmt"
	"sync"
)

func Cron() {
	fmt.Println("INFO: Starting data collection workers...")

	var wg sync.WaitGroup
	numWorkers := maxConcurrentFetch

	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("INFO: Starting worker %d\n", workerID)
			Worker(workerID)
		}(i)
	}

	
	wg.Wait()
	fmt.Println("INFO: All workers completed")
}
