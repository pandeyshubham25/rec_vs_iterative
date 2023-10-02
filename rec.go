package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	depth        = 10
	maxVariables = 50000
	step         = 2000
)

func recursiveTask(variables_count int, iterations int) {
	if iterations == 0 {
		return
	}
	// Initialize a slice with n elements
	variables := make([]int, variables_count)

	// Initialize variables in the slice
	for i := 0; i < variables_count; i++ {
		variables[i] = i
	}

	recursiveTask(variables_count, iterations-1)
}

func iterativeTask(variables_count int, iterations int) {
	for j := 0; j < iterations; j++ {
		variables := make([]int, variables_count)
		// Initialize variables in the slice
		for i := 0; i < variables_count; i++ {
			variables[i] = i
		}
	}
}

func measureTime(executionFunc func(int, int), variables_count int, depth int) time.Duration {
	start := time.Now()
	executionFunc(variables_count, depth)
	return time.Since(start)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for variables_count := 1; variables_count < maxVariables; variables_count += step {
		recursiveTotalTime := 0.0
		iterativeTotalTime := 0.0
		for j := 0; j < 5; j++ {
			recursiveTime := measureTime(recursiveTask, variables_count, depth).Seconds()
			iterativeTime := measureTime(iterativeTask, variables_count, depth).Seconds()
			recursiveTotalTime += recursiveTime
			iterativeTotalTime += iterativeTime

		}
		recursiveAvgTime := recursiveTotalTime / float64(5)
		iterativeAvgTime := iterativeTotalTime / float64(5)

		fmt.Printf("Variables: %d, Recursive Avg Time: %f seconds, Iterative Avg Time: %f seconds\n", variables_count, recursiveAvgTime, iterativeAvgTime)
	}

}
