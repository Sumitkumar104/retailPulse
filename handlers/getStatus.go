package handlers

import (
	"encoding/json"
	"net/http"
	"retailPulse/models"
)

func GetJobStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, `{"error":"Missing job ID"}`, http.StatusBadRequest)
		return
	}

	jobInterface, ok := models.Jobs.Load(jobID)
	if !ok {
		http.Error(w, `{"error":"Job not found"}`, http.StatusBadRequest)
		return
	}

	job := jobInterface.(*models.Job)
	json.NewEncoder(w).Encode(job)
}
