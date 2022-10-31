package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const FarmStatus = "Farming status:"
const PlotCount = "Plot count for all harvesters:"
const TotalPlotSize = "Total size of plots:"
const ExpectedTimeToWin = "Expected time to win:"
const EstimatedNetspace = "Estimated network space:"
const CurrentDifficulty = "Current difficulty:"
const Network = "Network:"
const TotalIterations = "Total iterations since the start of the blockchain:"
const CurrentBlockChainStatus = "Current Blockchain Status:" // If not synced local and global height is under this
const Height = "Height:"                                     // If synced
const TotalBalance = "-Total Balance:"

type TStats struct {
	Ettw            string // Expected time to win
	Netspace        string
	FarmStatus      string
	PlotCount       int // -1 = Unknown
	TotalPlotSize   string
	Difficulty      int // -1 = Unknown or able to parse
	Network         string
	TotalIterations int // Iterations since blockchain start
	LocalHeight     int
	GlobalHeight    int
	TotalBalance    float64
	LatestError     string
}

func RefreshStats() (TStats, error) {
	var stats TStats = TStats{}

	rawData, err := fetchDataFromCommands()
	if err != nil {
		return stats, err
	}

	parseCommandOutput(rawData, &stats)
	parseLogOutput(&stats)

	return stats, nil
}

func parseCommandOutput(rawData []string, stats *TStats) {
	// Get information from cmd output
	for _, line := range rawData {
		if strings.Contains(line, PlotCount) && stats.PlotCount == 0 {
			// Plot Count
			plotCount, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(line, PlotCount), " "))
			if err != nil {
				stats.PlotCount = -1
			} else {
				stats.PlotCount = plotCount
			}
		} else if strings.Contains(line, TotalPlotSize) && strings.EqualFold(stats.TotalPlotSize, "") {
			// Total size of plots
			stats.TotalPlotSize = strings.Trim(strings.TrimPrefix(line, TotalPlotSize), " ")
		} else if strings.Contains(line, FarmStatus) && strings.EqualFold(stats.FarmStatus, "") {
			// Farming Status
			stats.FarmStatus = strings.Trim(strings.TrimPrefix(line, FarmStatus), " ")
		} else if strings.Contains(line, ExpectedTimeToWin) && strings.EqualFold(stats.Ettw, "") {
			// Expected time to win
			stats.Ettw = strings.Trim(strings.TrimPrefix(line, ExpectedTimeToWin), " ")
		} else if strings.Contains(line, EstimatedNetspace) && strings.EqualFold(stats.Netspace, "") {
			// Netspace
			stats.Netspace = strings.Trim(strings.TrimPrefix(line, EstimatedNetspace), " ")
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
		} else if strings.Contains(line, TotalIterations) && stats.TotalIterations == 0 {
			// Total Iterations since blockchain start
			iterations, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(line, TotalIterations), " "))
			if err != nil {
				stats.TotalIterations = -1
			} else {
				stats.TotalIterations = iterations
			}
		} else if strings.Contains(line, CurrentBlockChainStatus) && stats.LocalHeight == 0 && stats.GlobalHeight == 0 {
			var tmp = strings.Split(strings.TrimPrefix(line, CurrentBlockChainStatus), " ")
			for _, entry := range tmp {
				if strings.Contains(entry, "/") {
					tmp = strings.Split(tmp[2], "/")
				}
			}

			localHeight, err := strconv.Atoi(tmp[0])
			if err != nil {
				stats.LocalHeight = -1
			} else {
				stats.LocalHeight = localHeight
			}
			globalHeight, err := strconv.Atoi(tmp[1])
			if err != nil {
				stats.GlobalHeight = -1
			} else {
				stats.GlobalHeight = globalHeight
			}
		} else if strings.Contains(line, Height) && stats.LocalHeight == -1 && stats.GlobalHeight == -1 {
			tmp, err := strconv.Atoi(strings.Trim(strings.Split(line, Height)[1], " "))
			if err == nil {
				stats.LocalHeight = tmp
				stats.GlobalHeight = tmp
			} // Ignore if failed. Just keep -1 as Heights
		} else if strings.Contains(line, TotalBalance) && stats.TotalBalance == 0 {
			tmp, err := strconv.ParseFloat(Unique(strings.Split(line, " "))[2], 64)
			if err != nil {
				log.Println("Failed to get total wallet balance")
				stats.TotalBalance = -1
			} else {
				stats.TotalBalance = tmp
			}
		}
	}
}

func fetchDataFromCommands() ([]string, error) {
	var rawData []string

	// Run "chia farm summary"
	data, err := exec.Command(config.ChiaPath, "farm", "summary").Output()
	if err != nil {
		log.Fatal("Failed to get farm summary")
		return nil, err
	}
	rawData = strings.Split(string(data), "\n")

	// Run "chia show -s"
	data, err = exec.Command(config.ChiaPath, "show", "-s").Output()
	if err != nil {
		log.Fatal("Failed to get current state of blockchain")
		return nil, err
	}
	rawData = append(rawData, strings.Split(string(data), "\n")...)

	// Run "chia wallet show"
	data, err = exec.Command(config.ChiaPath, "wallet", "show").Output()
	if err != nil {
		log.Fatal("Failed to get wallet info")
		return nil, err
	}
	rawData = append(rawData, strings.Split(string(data), "\n")...)

	return Unique(rawData), nil // Remove duplicated lines and return output data
}

func parseLogOutput(stats *TStats) {
	fileContent, err := ReadTextFile(config.ChiaLogPath)
	if err != nil {
		log.Fatal("Failed to read log file")
		return
	}

	for i := 0; i < len(fileContent); i++ {
		if strings.Contains(fileContent[i], "ERROR") {
			splitted := strings.Split(fileContent[i], "ERROR")
			stats.LatestError = strings.Trim(splitted[1], " ")
		}
	}
}
