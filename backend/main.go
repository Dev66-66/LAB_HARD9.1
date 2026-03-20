package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

// Task represents a range to calculate
type Task struct {
	Start int64
	End   int64
}

// Result represents the sum calculated for a range
type Result struct {
	Sum int64
}

// HeavyComputation calculates the sum of squares in a range using goroutines and channels.
func HeavyComputation(start, end int64, numWorkers int) int64 {
	if numWorkers <= 0 {
		numWorkers = runtime.NumCPU()
	}

	tasks := make(chan Task, numWorkers)
	results := make(chan int64, numWorkers)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				var partialSum int64
				for i := task.Start; i <= task.End; i++ {
					partialSum += i * i
				}
				results <- partialSum
			}
		}()
	}

	// Distribute work
	go func() {
		chunkSize := (end - start + 1) / int64(numWorkers)
		if chunkSize == 0 {
			chunkSize = end - start + 1
		}

		for i := start; i <= end; i += chunkSize {
			e := i + chunkSize - 1
			if e > end {
				e = end
			}
			tasks <- Task{Start: i, End: e}
		}
		close(tasks)
		wg.Wait()
		close(results)
	}()

	var totalSum int64
	for res := range results {
		totalSum += res
	}

	return totalSum
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")
	workersStr := r.URL.Query().Get("workers")

	start, _ := strconv.ParseInt(startStr, 10, 64)
	end, _ := strconv.ParseInt(endStr, 10, 64)
	workers, _ := strconv.Atoi(workersStr)

	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	result := HeavyComputation(start, end, workers)

	resp := map[string]interface{}{
		"result":  result,
		"start":   start,
		"end":     end,
		"workers": workers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/compute", computeHandler)
	port := ":8080"
	fmt.Printf("Go backend listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
