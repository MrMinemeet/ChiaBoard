package main

import (
	"fmt"
	"time"
)

var BuildVersion string = "NO VERSION SET. NOT BUILT PROPERLY?" // Set at compile time. Read from ./VERSION

// "github.com/leandroveronezi/go-terminal"

var config TConfig
var chiaVersion TVersion

func main() {
	Init()

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

func Init() {
	// Load Settings
	config = LoadSettings()
	chiaVersion, _ = GetChiaVersion()
}
