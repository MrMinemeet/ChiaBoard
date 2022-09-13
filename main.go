package main

import "time"

// "github.com/leandroveronezi/go-terminal"

func main() {

	for true {
		stats, err := RefreshStats()
		if err != nil {
			return
		}

		go PrintStats(stats)
		time.Sleep(2 * time.Second)
	}
}
