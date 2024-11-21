#include <opencv2/opencv.hpp>
#include <chrono>
#include <iostream>
#include <vector>
#include <cmath>
#include <numeric>

using namespace std;
using namespace cv;

Mat segmentCoins(const Mat &inputImage)
{
    // Convert to grayscale
    Mat gray;
    cvtColor(inputImage, gray, COLOR_BGR2GRAY);

    // Apply thresholding
    Mat thresh;
    threshold(gray, thresh, 100, 255, THRESH_BINARY_INV + THRESH_OTSU);

    // Morphological transformations for noise removal
    Mat kernel = getStructuringElement(MORPH_RECT, Size(3, 3));
    Mat opening;
    morphologyEx(thresh, opening, MORPH_OPEN, kernel, Point(-1, -1), 2);

    // Sure background area using dilation
    Mat sure_bg;
    dilate(opening, sure_bg, kernel, Point(-1, -1), 3);

    // Distance transform for sure foreground area
    Mat dist_transform;
    distanceTransform(opening, dist_transform, DIST_L2, 5);

    // Thresholding to find sure foreground
    Mat sure_fg;
    double maxVal;
    minMaxLoc(dist_transform, nullptr, &maxVal);
    threshold(dist_transform, sure_fg, 0.7 * maxVal, 255, 0);

    // Convert to 8-bit to find contours
    sure_fg.convertTo(sure_fg, CV_8U);
    vector<vector<Point>> contours;
    findContours(sure_fg, contours, RETR_EXTERNAL, CHAIN_APPROX_SIMPLE);

    // Draw contours on the original image
    Mat result = inputImage.clone();
    for (size_t i = 0; i < contours.size(); ++i)
    {
        drawContours(result, contours, static_cast<int>(i), Scalar(0, 255, 0), 2);
    }

    return result;
}

void benchmarkSegmentCoins(const Mat &image, int numRuns)
{
    vector<double> runTimes;

    for (int i = 0; i < numRuns; ++i)
    {
        auto start = chrono::high_resolution_clock::now();

        // Run the segmentation function
        Mat result = segmentCoins(image);

        auto end = chrono::high_resolution_clock::now();
        chrono::duration<double, micro> elapsed = end - start;

        runTimes.push_back(elapsed.count());
    }

    // Calculate min, max, mean, and standard deviation
    double minTime = *min_element(runTimes.begin(), runTimes.end());
    double maxTime = *max_element(runTimes.begin(), runTimes.end());
    double meanTime = accumulate(runTimes.begin(), runTimes.end(), 0.0) / runTimes.size();

    // Standard deviation calculation
    double variance = 0.0;
    for (const auto &time : runTimes)
    {
        variance += pow(time - meanTime, 2);
    }
    variance /= runTimes.size();
    double stdDev = sqrt(variance);

    // Output the benchmark results
    cout << "Benchmark Results (in microseconds):" << endl;
    cout << "Min Time: " << minTime << " µs" << endl;
    cout << "Max Time: " << maxTime << " µs" << endl;
    cout << "Mean Time: " << meanTime << " µs" << endl;
    cout << "Standard Deviation: " << stdDev << " µs" << endl;
}

int main()
{
    // Load the image from the assets folder
    Mat image = imread("../../water_coins.jpg");
    if (image.empty())
    {
        cerr << "Failed to load image." << endl;
        return -1;
    }

    // Benchmark the segmentCoins function
    int numRuns = 1000;
    benchmarkSegmentCoins(image, numRuns);

    return 0;
}
