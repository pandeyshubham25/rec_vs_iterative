package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	num_observations = 10
)

var (
	depth        int
	maxVariables int
	step         int
	fileName     string
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
	flag.IntVar(&depth, "depth", 10, "Depth value")
	flag.IntVar(&maxVariables, "maxVariables", 100000, "Max variables value")
	flag.IntVar(&step, "step", 5000, "Step value")
	flag.StringVar(&fileName, "fileName", "data.txt", "File Name")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Close the file when we're done
	for variables_count := 1; variables_count < maxVariables; variables_count += step {
		recursiveTotalTime := 0.0
		iterativeTotalTime := 0.0
		for j := 0; j < num_observations; j++ {
			recursiveTime := measureTime(recursiveTask, variables_count, depth).Seconds()
			iterativeTime := measureTime(iterativeTask, variables_count, depth).Seconds()
			recursiveTotalTime += recursiveTime
			iterativeTotalTime += iterativeTime

		}
		recursiveAvgTime := recursiveTotalTime / float64(num_observations)
		iterativeAvgTime := iterativeTotalTime / float64(num_observations)

		formattedString := fmt.Sprintf("Variables: %d, Recursive Avg Time: %f seconds, Iterative Avg Time: %f seconds\n", variables_count, recursiveAvgTime, iterativeAvgTime)
		_, err = fmt.Fprint(file, formattedString)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

}
