package main

import (
	"fmt"
	"time"
)

func PrintStats(stats TStats) {
	fmt.Println("        ", time.Now().Format("02.01.2006 - 15:04:05"))
	fmt.Println("=======================================")
	fmt.Println("Expected time to win:", stats.Ettw)
	fmt.Println("Estimated Netspace:", stats.Netspace)
	fmt.Println("Farm Status:", stats.FarmStatus)
	fmt.Println("Total Plot Count:", stats.PlotCount)
	fmt.Println("Difficulty:", stats.Difficulty)
	fmt.Println("Network:", stats.Network)
	fmt.Println("Iterations:", stats.TotalIterations)
	fmt.Println("=======================================")
}
