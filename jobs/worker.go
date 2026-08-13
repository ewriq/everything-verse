package jobs

import (
	"fmt"

	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup


func Data() {
	DataCollectionMu.Lock()

	if DataCollectionBusy {
		DataCollectionMu.Unlock()
		fmt.Println("INFO: running")
		return
	}

	DataCollectionBusy = true
	DataCollectionMu.Unlock()

	defer func() {
		DataCollectionMu.Lock()
		DataCollectionBusy = false
		DataCollectionMu.Unlock()
	}()

	sources, err := LoadSources(SourcesPath)
	if err != nil {
		fmt.Printf("ERROR: opml source fetch failed: %v\n", err)
		return
	}


	jobs := make(chan Source)

	workerCount := maxWorker


	wg.Add(workerCount)
	
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for src := range jobs {
				//fmt.Printf("INFO: fwetch data %s...\n", src.Name)
				inserted, err := ProcessSource(src)
				if err != nil {
					fmt.Printf("ERROR: fail to process %s: %v\n", src.Name, err)
					continue
				}
				if inserted {
					atomic.AddInt32(&InsertedCount, 1)
					//fmt.Printf("INFO: added data %s\n", src.Name)
				}
			}
		}()
	}

	for _, src := range sources {
		jobs <- src
	}

	close(jobs);

	wg.Wait();

	fmt.Printf("INFO: complete for %d sources\n", len(sources))
	if InsertedCount > 0 {
		fmt.Printf("INFO: successfully inserted from %d sources\n", InsertedCount)
	}
	
}


