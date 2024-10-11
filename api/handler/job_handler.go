package handler

import (
	"encoding/json"
	"net/http"

	"event-driven/api/usecase"
	"event-driven/common/model"

	"github.com/go-chi/chi/v5"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(jobUsecase usecase.JobUsecase) *JobHandler {
	return &JobHandler{jobUsecase: jobUsecase}
}

// Name: Handle Job
// Endpoint: /job
// Method: POST
// RequiresAuth: false
func (h *JobHandler) HandleJob(w http.ResponseWriter, r *http.Request) {
	var job model.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	jobID, err := h.jobUsecase.EnqueueJob(job)
	if err != nil {
		http.Error(w, "Failed to add job to queue", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"jobId": jobID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Name: Get Job Status
// Endpoint: /job/status/{jobID}
// Method: GET
// RequiresAuth: false
func (h *JobHandler) HandleJobStatus(w http.ResponseWriter, r *http.Request) {
	jobID := chi.URLParam(r, "jobID")
	status, err := h.jobUsecase.GetJobStatus(jobID)
	if err != nil {
		http.Error(w, "Failed to get job status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
