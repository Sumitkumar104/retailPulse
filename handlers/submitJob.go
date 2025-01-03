package handlers

import (
	"encoding/json"
	"net/http"
	"retailPulse/models"
	"retailPulse/utils"
	"sync"
	"time"
)

type SubmitRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

func SubmitJob(w http.ResponseWriter, r *http.Request) {
	var req SubmitRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Count != len(req.Visits) {
		http.Error(w, `{"error":"Invalid request"}`, http.StatusBadRequest)
		return
	}

	// Create Job
	job := models.Job{
		ID:        utils.GenerateJobID(),
		Status:    "ongoing",
		CreatedAt: time.Now(),
		Results:   make(map[string][]models.ImageResult),
	}

	// Process Job in Goroutine
	go processJob(&job, req.Visits)

	models.Jobs.Store(job.ID, &job)

	// Respond with Job ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"job_id": job.ID})
}

func processJob(job *models.Job, visits []Visit) {
	var wg sync.WaitGroup

	for _, visit := range visits {
		wg.Add(1)
		go func(visit Visit) {
			defer wg.Done()

			store, exists := models.StoreMaster[visit.StoreID]
			if !exists {
				job.StoreErrors = append(job.StoreErrors, models.StoreError{
					StoreID: visit.StoreID,
					Error:   "Store ID not found",
				})
				return
			}

			for _, imageURL := range visit.ImageURLs {
				perimeter, err := utils.ProcessImage(imageURL)
				if err != nil {
					job.StoreErrors = append(job.StoreErrors, models.StoreError{
						StoreID: visit.StoreID,
						Error:   "Failed to process image",
					})
					return
				}
				job.Results[store.ID] = append(job.Results[store.ID], models.ImageResult{
					ImageURL:  imageURL,
					Perimeter: perimeter,
				})
			}
		}(visit)
	}
	wg.Wait()

	if len(job.StoreErrors) > 0 {
		job.Status = "failed"
	} else {
		job.Status = "completed"
	}
}
