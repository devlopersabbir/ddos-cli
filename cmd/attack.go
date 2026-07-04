package cmd

import (
	"sync"
	"time"
)

func Attack(rps int64) {
	interval := time.Second / time.Duration(rps)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for range ticker.C {
		wg.Add(1)
	}

}
