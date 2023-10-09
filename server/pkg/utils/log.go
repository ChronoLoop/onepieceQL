package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

func PrintProgress(currentIteration, totalIterations, lastPrintedPercent int) int {
	percent := (currentIteration + 1) * 100 / totalIterations

	if percent != lastPrintedPercent {
		fmt.Printf("Progress: %d%%\n", percent)
		lastPrintedPercent = percent
	}

	return lastPrintedPercent
}

func PrintStruct(s interface{}) {
	bytes, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("\nstruct: %s\n", string(bytes))
}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
