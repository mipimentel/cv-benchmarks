package main

import (
	"fmt"
	"image" // Import the standard library package for creating points (Pt)
	"image/color"
	"math"
	"time"

	"gocv.io/x/gocv"
)

func segmentCoins(inputImage gocv.Mat) (gocv.Mat, int, []float64) {
	// Convert to grayscale
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(inputImage, &gray, gocv.ColorBGRToGray)

	// Apply thresholding
	thresh := gocv.NewMat()
	defer thresh.Close()
	gocv.Threshold(gray, &thresh, 100, 255, gocv.ThresholdBinaryInv+gocv.ThresholdOtsu)

	// Noise removal with Morphological Transformations
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3)) // Corrected usage with image.Pt
	defer kernel.Close()
	opening := gocv.NewMat()
	defer opening.Close()
	gocv.MorphologyEx(thresh, &opening, gocv.MorphOpen, kernel)

	// Sure background area using dilation
	sureBg := gocv.NewMat()
	defer sureBg.Close()
	gocv.Dilate(opening, &sureBg, kernel)

	// Finding sure foreground area using Distance Transform
	distTransform := gocv.NewMat()
	defer distTransform.Close()
	labels := gocv.NewMat()
	defer labels.Close()
	gocv.DistanceTransform(opening, &distTransform, &labels, gocv.DistL2, gocv.DistanceMask5, gocv.DistanceLabelCComp)
	_, maxVal, _, _ := gocv.MinMaxLoc(distTransform)

	sureFg := gocv.NewMat()
	defer sureFg.Close()
	gocv.Threshold(distTransform, &sureFg, 0.7*maxVal, 255, gocv.ThresholdBinary)

	// Finding contours from the cleaned binary image
	sureFg.ConvertTo(&sureFg, gocv.MatTypeCV8U)
	contours := gocv.FindContours(sureFg, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	// Drawing contours and calculating areas
	processedImage := inputImage.Clone()
	defer processedImage.Close()
	coinAreas := []float64{}
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		area := gocv.ContourArea(contour)
		coinAreas = append(coinAreas, area)
		gocv.DrawContours(&processedImage, contours, i, color.RGBA{0, 255, 0, 0}, 2)
	}

	return processedImage, contours.Size(), coinAreas
}

func main() {
	// Number of times to run the benchmark
	numRuns := 1000
	runTimes := make([]float64, numRuns)

	// Load the image from the root directory
	imagePath := "../water_coins.jpg"
	image := gocv.IMRead(imagePath, gocv.IMReadColor)
	if image.Empty() {
		panic("Failed to read image: " + imagePath)
	}
	defer image.Close()

	// Run the segmentCoins function multiple times
	for i := 0; i < numRuns; i++ {
		start := time.Now()
		_, _, _ = segmentCoins(image)
		elapsed := time.Since(start).Microseconds() // Capture elapsed time in microseconds
		runTimes[i] = float64(elapsed)
	}

	// Calculate min, max, mean, and standard deviation
	minTime, maxTime, meanTime, stdDev := calculateStatistics(runTimes)

	// Print out the results
	fmt.Printf("Benchmark Results (in microseconds):\n")
	fmt.Printf("Min Time: %.2f µs\n", minTime)
	fmt.Printf("Max Time: %.2f µs\n", maxTime)
	fmt.Printf("Mean Time: %.2f µs\n", meanTime)
	fmt.Printf("Standard Deviation: %.2f µs\n", stdDev)
}

func calculateStatistics(times []float64) (min, max, mean, stdDev float64) {
	if len(times) == 0 {
		return 0, 0, 0, 0
	}

	// Calculate min, max, and mean
	min, max = times[0], times[0]
	sum := 0.0
	for _, t := range times {
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
		sum += t
	}
	mean = sum / float64(len(times))

	// Calculate standard deviation
	varianceSum := 0.0
	for _, t := range times {
		varianceSum += math.Pow(t-mean, 2)
	}
	variance := varianceSum / float64(len(times))
	stdDev = math.Sqrt(variance)

	return min, max, mean, stdDev
}
