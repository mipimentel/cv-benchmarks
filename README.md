# Computer Vision Benchmark

This is a repo I created to compare the performance of Python and Go, driven by the common hypothesis that simply changing to a compiled language like Go would automatically lead to faster execution times. Many people suggest, "Why not use Go for a performance boost, since it’s compiled?" I wanted to evaluate how significant these gains would actually be and how troublesome the implementation might become.

So, how much performance improvement can we really expect, and is it worth the effort to rewrite the codebase?

## The Test

I used a coin segmentation example from OpenCV (the watershed example) and employed LLMs to generate the code. The code may not be fully optimized but serves as a "quick and dirty" benchmark to answer whether Python's perceived "slowness" is truly a bottleneck for computer vision tasks.

## Benchmark Overview

In this benchmark, I conducted benchmarks in Go and Python to compare the performance of image segmentation algorithms for detecting and counting coins. The benchmarks were run using OpenCV in both languages to ensure consistent functionality.

### Test Environment

All benchmarks were run on the following hardware and software configurations:

- **CPU**: AMD Ryzen 5 5600X 6-Core Processor
- **RAM**: 32 GB DDR4
- **OS**: Ubuntu 24.04 LTS
- **Python Version**: 3.12
- **Go Version**: 1.23
- **OpenCV Version**: 4.10.0

### Benchmarking Methodology

The `segmentCoins()` function was benchmarked in both Go and Python. This function takes an image of coins and segments, counts, and finds the areas of individual coins. The benchmark was run with the `water_coins.jpg` test image, and the following metrics were collected:

- **Time per Execution** (`ns/op` or `µs/op`)
- **Memory Allocations per Operation** (`allocs/op`)
- **Total Memory Used** (`B/op`)

### Benchmark Results

| Language        | Min Time (µs) | Max Time (µs) | Mean Time (µs) | StdDev (µs) | Memory Allocations (B/op) | Allocations per Op |
|-----------------|---------------|---------------|----------------|-------------|--------------------------|--------------------|
| **Go (GoCV)**   | 798.00        | 4034.00       | 1226.44        | 176.06      | 536                      | 10                 |
| **Python (cv2)**| 714.60        | 1068.37       | 770.65         | 83.77       | -                        | -                  |

### Observations:

- **Python Execution Time**: The Python implementation had a faster mean execution time of **770.65 µs**, compared to Go's **1226.44 µs**.
- **Go Performance Variance**: Go had a larger **maximum execution time** (**4034 µs**) and a higher **standard deviation** (**176.06 µs**), suggesting more performance variability.
- **Memory Allocations**: Go's implementation shows **10 memory allocations per operation** on average, possibly due to garbage collection and extra memory overhead. Python’s C++ bindings handle memory differently, and pytest-benchmark may not capture the full picture, so take this result with a grain of salt.

### Analysis and Insights

- Different from the expected behavior, using go instead of Python for image processing tasks did not result in **faster execution times** just by changing the language.
- Go's garbage collection and memory management may not be fully optimized, compared to Python's C++ bindings, but this need more investigation, due to the lack of information from pytest-benchmark. This may explain the larger variance in execution times, as seen in the **higher standard deviation**.
- Optimizing Go’s implementation by reducing memory allocations could potentially improve its performance, but it needs more expertise of the developer compared to Python implementations.

### Conclusion

While Go is often seen as a performant alternative to Python, our benchmarks reveal that, for OpenCV-based image processing, Python’s mature bindings and efficient C++ backend provide superior performance in terms of both speed and stability. However, Go still is a faster language comparered to Python if Python is not using C++ bindings.

From this quick implementation and set of tests, the hypothesis that simply changing the language would result in faster execution times was not confirmed. Changing the entire codebase to a different language may not be worth the effort due to the increased complexity in implementation and higher development cost. If execution performance is the primary focus, C++ might be a better choice, while Python remains advantageous for rapid development and prototyping due to its ease of use and mature ecosystem.
