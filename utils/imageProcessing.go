package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"time"
)

func ProcessImage(url string) (int, error) {   // return perimeter of image
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return 0, err
	}

	// Calculate perimeter
	perimeter := 2 * (img.Height + img.Width)

	// Simulate GPU processing
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

	return perimeter, nil
}
