// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

func main() {
	// Set up server routes
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)

	// Start the server
	port := 8000
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func processSingle(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var req SortRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Perform sequential sorting
	start := time.Now()
	time.Sleep(5 * time.Second)
	sortedArrays := make([][]int, len(req.ToSort))
	for i, arr := range req.ToSort {
		sortedArrays[i] = make([]int, len(arr))
		copy(sortedArrays[i], arr)
		sort.Ints(sortedArrays[i])
	}
	timeTaken := time.Since(start).Round(time.Nanosecond).Nanoseconds()

	// Prepare and send JSON response
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       timeTaken,
	}
	sendJSONResponse(w, response)
}

func processConcurrent(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var req SortRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Perform concurrent sorting
	start := time.Now()
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	var mu sync.Mutex
	sortedArrays := make([][]int, len(req.ToSort))
	for i, arr := range req.ToSort {
		wg.Add(1)
		go func(i int, arr []int) {
			defer wg.Done()
			sorted := make([]int, len(arr))
			copy(sorted, arr)
			sort.Ints(sorted)
			mu.Lock()
			sortedArrays[i] = sorted
			mu.Unlock()
		}(i, arr)
	}
	wg.Wait()
	elapsed := time.Since(start)
	timeTaken := elapsed.Nanoseconds()

	// Prepare and send JSON response
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       timeTaken,
	}
	sendJSONResponse(w, response)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}
