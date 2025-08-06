package jobs

import (
	
	"fmt"

	"time"
)

func Worker(id int) {
	for {
		fmt.Println("[Start", id, "]")
		Data()
		fmt.Println("[End", id, "]")
		time.Sleep(24 * time.Second) 
	}
}

