package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
    "time"
    "github.com/gorilla/mux"
)

var jobIDCounter int
var jobs = sync.Map{}

// JobRequest represents the structure of the incoming job submission
type JobRequest struct {
    Count  int `json:"count"`
    Visits []struct {
        StoreID   string   `json:"store_id"`
        ImageURLs []string `json:"image_url"`
        VisitTime string   `json:"visit_time"`
    } `json:"visits"`
}

// JobResponse represents the response returned after a job is created
type JobResponse struct {
    JobID int `json:"job_id"`
}

// JobStatus represents the status of a job
type JobStatus struct {
    Status string `json:"status"`
    JobID  int    `json:"job_id"`
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/submit/", SubmitJobHandler).Methods("POST")
    router.HandleFunc("/api/status", GetJobStatusHandler).Methods("GET")

    log.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", router)
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
    var jobRequest JobRequest
    err := json.NewDecoder(r.Body).Decode(&jobRequest)
    if err != nil || jobRequest.Count != len(jobRequest.Visits) {
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }

    // Generate a new job ID
    jobIDCounter++
    jobID := jobIDCounter

    // Start processing the job asynchronously
    go processJob(jobID, jobRequest)

    // Return the job ID as response
    response := JobResponse{JobID: jobID}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func processJob(jobID int, jobRequest JobRequest) {
    for _, visit := range jobRequest.Visits {
        processImages(visit.ImageURLs)
    }

    // Mark job as completed (for simplicity, using in-memory sync.Map)
    jobs.Store(jobID, "completed")
    log.Printf("Job %d completed", jobID)
}

func processImages(imageURLs []string) {
    for _, imageURL := range imageURLs {
        // Simulate image processing: calculate perimeter and random sleep
        height, width := 100, 200 // Example dimensions
        perimeter := 2 * (height + width)
        log.Printf("Processing image %s: Perimeter = %d", imageURL, perimeter)

        // Random sleep to simulate processing time
        randomSleep := time.Duration(rand.Intn(300)+100) * time.Millisecond
        time.Sleep(randomSleep)
    }
}

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
    jobIDStr := r.URL.Query().Get("jobid")
    jobID, err := strconv.Atoi(jobIDStr)
    if err != nil {
        http.Error(w, "Invalid job ID", http.StatusBadRequest)
        return
    }

    // Check job status
    status, ok := jobs.Load(jobID)
    if !ok {
        http.Error(w, "Job ID not found", http.StatusBadRequest)
        return
    }

    response := JobStatus{Status: status.(string), JobID: jobID}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
