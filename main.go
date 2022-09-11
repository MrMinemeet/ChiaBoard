package main

import "log"

// "github.com/leandroveronezi/go-terminal"

func main() {

	stats, err := RefreshStats()
	if err != nil {
		return
	}

	log.Println("Expected time to win:", stats.Ettw)
	log.Println("Estimated Netspace:", stats.Netspace)
	log.Println("Farm Status:", stats.FarmStatus)
	log.Println("Plot Count:", stats.PlotCount)
	log.Println("Diffculty:", stats.Difficulty)
	log.Println("Network:", stats.Network)
}
