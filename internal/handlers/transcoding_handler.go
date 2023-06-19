package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"time"

	"github.com/gorilla/mux"

	"github.com/perkdrew/media-transcoder/internal/models"
	"github.com/perkdrew/media-transcoder/internal/transcoder"
)

type JobIDGenerator interface {
	Generate() string
}

// RandomJobIDGenerator generates a random job ID
type RandomJobIDGenerator struct{}

// Generate generates a random job ID
func (g *RandomJobIDGenerator) Generate() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// TranscodingHandler handles the transcoding API requests
type TranscodingHandler struct {
	Transcoder     *transcoder.Transcoder
	Jobs           map[string]*models.Job
	JobIDGenerator JobIDGenerator
}

// CreateTranscodeJob handles the creation of a new transcoding job
func (h *TranscodingHandler) CreateTranscodeJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique job ID
	jobID := h.JobIDGenerator.Generate()

	// Create a new Job instance
	newJob := models.NewJob(jobID, job.InputFile, job.OutputFile, job.Parameters) // Added job.Parameters

	// Add the job to the jobs map
	h.Jobs[jobID] = newJob

	// Perform transcoding in a separate goroutine
	go h.performTranscoding(newJob)

	// Send the job ID as the response
	response := struct {
		JobID string `json:"jobID"`
	}{
		JobID: jobID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


// GetJobStatus retrieves the status and details of a transcoding job
func (h *TranscodingHandler) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jobID := params["jobID"]

	job, ok := h.Jobs[jobID]
	if !ok {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// Helper function to perform transcoding in a separate goroutine
func (h *TranscodingHandler) performTranscoding(job *models.Job) {
	err := h.Transcoder.Transcode(job.InputFile, job.OutputFile, job.Parameters) // use the job parameters here
	if err != nil {
		log.Printf("Transcoding failed for job %s: %s", job.ID, err)
		job.Status = "failed"
		return
	}

	job.Status = "completed"
	job.Progress = 100
}

// Helper function to start transcoding in a separate goroutine
func (h *TranscodingHandler) startTranscoding(job *models.Job) {
	h.performTranscoding(job)
}

// NewTranscodingHandler creates a new TranscodingHandler instance
func NewTranscodingHandler(transcoder *transcoder.Transcoder, jobIDGenerator JobIDGenerator) *TranscodingHandler {
	return &TranscodingHandler{
		Transcoder:     transcoder,
		Jobs:           make(map[string]*models.Job),
		JobIDGenerator: jobIDGenerator,
	}
}

