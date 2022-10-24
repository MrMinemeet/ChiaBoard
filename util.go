package main

import (
	"fmt"
	"log"
	"os/exec"
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

func GetChiaVersion() (TVersion, error) {
	// Execute "chia version"
	data, err := exec.Command(config.ChiaPath, "version").Output()
	if err != nil {
		log.Fatal("Failed to execute \"chia version\"")
		return TVersion{}, err
	}

	version, err := GetVersion(string(data))
	if err != nil {
		log.Fatal("Failed to parse chia version")
		return TVersion{}, err
	}

	PrintVersion(version)

	return version, nil
}
