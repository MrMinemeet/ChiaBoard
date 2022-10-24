package main

import (
	"fmt"
	"log"
	"os"
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
	fmt.Println("Height:", stats.LocalHeight, "/", stats.GlobalHeight)
	fmt.Println("Total Balance:", stats.TotalBalance, "XCH")
	fmt.Println("Latest Error:", stats.LatestError)
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

func ReadTextFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Cannot open log file.", err)
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
