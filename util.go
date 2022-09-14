package main

import (
	"fmt"
	"strings"
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

func Unique(data []string) []string {
	var contains = make(map[string]bool)
	var result []string

	for _, entry := range data {
		if _, contained := contains[entry]; !contained && !strings.EqualFold(entry, "") {
			// Check if string was not added already and string is not empty
			contains[entry] = true
			result = append(result, entry)
		}
	}

	return result
}
