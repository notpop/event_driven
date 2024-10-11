package handler

import (
	"encoding/json"
	"net/http"

	"event-driven/api/websocket"
)

type JobStatusHandler struct{}

func NewJobStatusHandler() *JobStatusHandler {
	return &JobStatusHandler{}
}

// Name: Update Job Status
// Endpoint: /job/status/update
// Method: POST
// RequiresAuth: false
func (h *JobStatusHandler) UpdateJobStatus(w http.ResponseWriter, r *http.Request) {
	var msg map[string]string
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	websocket.BroadcastMessage(websocket.Message{
		JobID:  msg["jobId"],
		Status: msg["status"],
	})

	w.WriteHeader(http.StatusOK)
}
