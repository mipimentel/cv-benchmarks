// File: project-root/go-cv/main_test.go
package main

import (
	"testing"

	"gocv.io/x/gocv"
)

func BenchmarkSegmentCoins(b *testing.B) {
	// Load the image from the root directory
	imagePath := "../water_coins.jpg"
	image := gocv.IMRead(imagePath, gocv.IMReadColor)
	if image.Empty() {
		b.Fatalf("Failed to read image: %s\n", imagePath)
	}

	// Reset the timer so that the load time isn't counted in the benchmark
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = segmentCoins(image)
	}
}
