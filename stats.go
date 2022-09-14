package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const FarmStatus = "Farming status:"
const PlotCount = "Plot count for all harvesters:"
const ExpectedTimeToWin = "Expected time to win:"
const EstimatedNetspace = "Estimated network space:"
const CurrentDifficulty = "Current difficulty:"
const Network = "Network:"

type TStats struct {
	Ettw       string // Expected time to win
	Netspace   string
	FarmStatus string
	PlotCount  int // -1 = Unknown
	Difficulty int // -1 = Unknown or able to parse
	Network    string
}

func RefreshStats() (TStats, error) {
	var stats TStats = TStats{}
	var output []string

	// Run "chia farm summary"
	data, err := exec.Command(TmpChiaPath, "farm", "summary").Output()
	if err != nil {
		log.Fatal("Failed to get farm summary")
		return stats, err
	}
	output = strings.Split(string(data), "\n")

	// Run "chia show -s"
	data, err = exec.Command(TmpChiaPath, "show", "-s").Output()
	if err != nil {
		log.Fatal("Failed to get current state of blockchain")
		return stats, err
	}
	output = append(output, strings.Split(string(data), "\n")...)

	// Get information from cmd output
	for _, line := range output {
		if strings.Contains(line, PlotCount) && stats.PlotCount == 0 {
			// Plot Count
			plotCount, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(line, PlotCount), " "))
			if err != nil {
				stats.PlotCount = -1
			} else {
				stats.PlotCount = plotCount
			}
		} else if strings.Contains(line, FarmStatus) && strings.EqualFold(stats.FarmStatus, "") {
			// Farming Status
			stats.FarmStatus = strings.TrimPrefix(line, FarmStatus)
		} else if strings.Contains(line, ExpectedTimeToWin) && strings.EqualFold(stats.Ettw, "") {
			// Expected time to win
			stats.Ettw = strings.TrimPrefix(line, ExpectedTimeToWin)
		} else if strings.Contains(line, EstimatedNetspace) && strings.EqualFold(stats.Netspace, "") {
			// Netspace
			stats.Netspace = strings.TrimPrefix(line, EstimatedNetspace)
		} else if strings.Contains(line, CurrentDifficulty) && stats.Difficulty == 0 {
			// Difficulty
			difficulty, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(line, CurrentDifficulty), " "))
			if err != nil {
				stats.Difficulty = -1
			} else {
				stats.Difficulty = difficulty
			}
		} else if strings.Contains(line, Network) && strings.EqualFold(stats.Network, "") {
			// Network
			tmp := strings.Split(line, " ")
			stats.Network = tmp[1]
		}
	}

	return stats, nil
}
