package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TVersion struct {
	Major int
	Minor int
	Patch string
}

func PrintVersion(v TVersion) {
	// Print version split with dot but without spaces
	fmt.Printf("Chia Version: %d.%d.%s", v.Major, v.Minor, v.Patch)
}

func GetVersion(v string) (TVersion, error) {
	var version TVersion = TVersion{}
	var splitted []string = strings.Split(v, ".")

	if len(splitted) != 3 {
		log.Println("Failed to parse version string")
		return version, errors.New("invalid version format")
	}

	// Major
	val, err := strconv.Atoi(splitted[0])
	if err != nil {
		return TVersion{}, err
	} else {
		version.Major = val
	}

	// Minor
	val, err = strconv.Atoi(splitted[1])
	if err != nil {
		return TVersion{}, err
	} else {
		version.Minor = val
	}

	// Patch
	version.Patch = splitted[2]

	return version, nil
}
