package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateJob(t *testing.T) {
	payload := []byte(`{"data": "test data"}`)
	req, _ := http.NewRequest("POST", "/job/create", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createJobHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var job Job
	json.Unmarshal(rr.Body.Bytes(), &job)
	if job.Data != "test data" || job.Status != "pending" {
		t.Errorf("job created incorrectly: %+v", job)
	}
}

func TestUpdateJob(t *testing.T) {
	// Pre-populate a job
	jobsMu.Lock()
	jobs[1] = &Job{ID: 1, Data: "raw", Status: "pending"}
	jobsMu.Unlock()

	payload := []byte(`{"id": 1, "encrypted": "encoded"}`)
	req, _ := http.NewRequest("PATCH", "/job/update", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateJobHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	if jobs[1].Encrypted != "encoded" || jobs[1].Status != "completed" {
		t.Errorf("job update failed: %+v", jobs[1])
	}
}
