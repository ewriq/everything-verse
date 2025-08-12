package jobs

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int) {
	for {
		fmt.Printf("INFO: Worker %d starting data collection\n", id)
		start := time.Now()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			Data()
		}()

		go func() {
			defer wg.Done()
			DeepSearch()
		}()

		go func() {
			defer wg.Done()
			Crawling()
		}()


		wg.Wait()

		duration := time.Since(start)
		fmt.Printf("INFO: Worker %d completed in %v\n", id, duration)
	}
}
