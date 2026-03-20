package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Job struct {
	ID        int    `json:"id"`
	Data      string `json:"data"`
	Encrypted string `json:"encrypted,omitempty"`
	Status    string `json:"status"` // "pending", "completed"
}

var (
	jobs    = make(map[int]*Job)
	nextID  = 1
	jobsMu  sync.Mutex
)

func createJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobsMu.Lock()
	job := &Job{
		ID:     nextID,
		Data:   payload.Data,
		Status: "pending",
	}
	jobs[nextID] = job
	nextID++
	jobsMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func getJobHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	jobsMu.Lock()
	job, ok := jobs[id]
	jobsMu.Unlock()

	if !ok {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func updateJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		ID        int    `json:"id"`
		Encrypted string `json:"encrypted"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobsMu.Lock()
	job, ok := jobs[payload.ID]
	if ok {
		job.Encrypted = payload.Encrypted
		job.Status = "completed"
	}
	jobsMu.Unlock()

	if !ok {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/job/create", createJobHandler)
	http.HandleFunc("/job/status", getJobHandler)
	http.HandleFunc("/job/update", updateJobHandler)

	port := ":8081"
	fmt.Printf("Orchestrator listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
