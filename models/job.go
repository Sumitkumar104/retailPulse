package models

import (
	"sync"
	"time"
)

type Job struct {
	ID          string
	Status      string
	StoreErrors []StoreError
	Results     map[string][]ImageResult
	CreatedAt   time.Time
}

type StoreError struct {
	StoreID string
	Error   string
}

type ImageResult struct {
	ImageURL string
	Perimeter int
}

var Jobs = sync.Map{} // Thread-safe map to store jobs or avoid Race condition
