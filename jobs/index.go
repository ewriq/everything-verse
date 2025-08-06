package jobs

import "sync"

func Cron() {
	var wg sync.WaitGroup

	numWorkers := 2
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			Worker(workerID)
		}(i)
	}

	wg.Wait()
}