package main

import (
	"fmt"
	"time"
)

var BuildVersion string // Set at compile time. Read from ./VERSION

// "github.com/leandroveronezi/go-terminal"

func main() {
	fmt.Println("Build Version:", BuildVersion)
	time.Sleep(500 * time.Millisecond)

	for true {
		stats, err := RefreshStats()
		if err != nil {
			return
		}

		go PrintStats(stats)
		time.Sleep(2 * time.Second)
	}
}
