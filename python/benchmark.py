import cv2
import numpy as np
import pytest


base_dir = __file__.rsplit("/", 1)[0].rsplit("/", 1)[0]
image_path = f"{base_dir}/water_coins.jpg"


def segment_coins(image):
    """
    Function to process the coin image using thresholding, noise removal, and contour detection.
    """
    # Step 1: Load the image

    # Convert to grayscale
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    # Step 2: Apply thresholding
    _, thresh = cv2.threshold(gray, 100, 255, cv2.THRESH_BINARY_INV + cv2.THRESH_OTSU)

    # Step 3: Noise removal with Morphological Transformations
    kernel = np.ones((3, 3), np.uint8)

    # Removing noise using morphological opening
    opening = cv2.morphologyEx(thresh, cv2.MORPH_OPEN, kernel, iterations=2)

    # Sure background area
    sure_bg = cv2.dilate(opening, kernel, iterations=3)

    # Finding sure foreground area
    dist_transform = cv2.distanceTransform(opening, cv2.DIST_L2, 5)
    _, sure_fg = cv2.threshold(dist_transform, 0.7 * dist_transform.max(), 255, 0)

    # Finding unknown region
    sure_fg = np.uint8(sure_fg)

    # Step 4: Find contours from the cleaned binary image
    contours, _ = cv2.findContours(sure_fg, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

    # Step 5: Draw contours and calculate areas
    processed_image = image.copy()
    coin_areas = []
    for contour in contours:
        area = cv2.contourArea(contour)
        coin_areas.append(area)
        cv2.drawContours(processed_image, [contour], -1, (0, 255, 0), 2)

    return processed_image, len(contours), coin_areas


@pytest.mark.benchmark
def test_coin_segmentation(benchmark):
    """
    Benchmark test for coin segmentation pipeline.
    """
    image = cv2.imread(image_path)

    # Benchmark the segmentation process
    result = benchmark(segment_coins, image)

    # Assert the number of detected coins
    assert 24 == result[1]

    # Optionally visualize the processed image (comment this in pytest runs)
    cv2.imshow("Segmented Coins", result[0])
    cv2.waitKey(0)
    cv2.destroyAllWindows()
