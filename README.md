# Computer Vision Benchmark

This is a repo I created to compare the performance of Python and Go, driven by the common hypothesis that simply changing to a compiled language like Go would automatically lead to faster execution times. Many people suggest, "Why not use Go for a performance boost, since it’s compiled?" I wanted to evaluate how significant these gains would actually be and how troublesome the implementation might become.

So, how much performance improvement can we really expect, and is it worth the effort to rewrite the codebase?

## The Test

I used a coin segmentation example from OpenCV (the watershed example) and employed LLMs to generate the code. The code may not be fully optimized but serves as a "quick and dirty" benchmark to answer whether Python's perceived "slowness" is truly a bottleneck for computer vision tasks.

Later was added a C++ implementation of the same code for sanity checks, since the go did not work as expected.

## Benchmark Overview

In this benchmark, I conducted benchmarks in Go, Python and C++ to compare the performance of image segmentation algorithms for detecting and counting coins. The benchmarks were run using OpenCV in both languages to ensure consistent functionality.

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
| **C++ (OpenCV)**| 402.77        | 1189.77       | 424.73         | 48.14       | -                        | -                  |

### Observations:

- **Python Execution Time**: The Python implementation had a faster mean execution time of **770.65 µs**, compared to Go's **1226.44 µs**.
- **Go Performance Variance**: Go had a larger **maximum execution time** (**4034 µs**) and a higher **standard deviation** (**176.06 µs**), suggesting more performance variability.
- **Memory Allocations**: Go's implementation shows **10 memory allocations per operation** on average, possibly due to garbage collection and extra memory overhead. Python’s C++ bindings handle memory differently, and pytest-benchmark may not capture the full picture, so take this result with a grain of salt.
- **C++ Execution Time**: The C++ version showed significant improvement with a **minimum time** of **402.77 µs** and a **mean time** of **424.73 µs**, which is the fastest among all three languages. This further highlights C++'s efficiency for such low-level tasks.

### Analysis and Insights

- Contrary to the expected behavior of compiled languages like go, using go instead of Python for image processing tasks did not result in **faster execution times** just by changing the language.
- Go's garbage collection and memory management may not be fully optimized, compared to Python's C++ bindings, but this need more investigation, due to the lack of information from pytest-benchmark. This may explain the larger variance in execution times, as seen in the **higher standard deviation**.
- Optimizing Go’s implementation by reducing memory allocations could potentially improve its performance, but it needs more expertise of the developer compared to Python implementations.
- **C++ Efficiency**: The C++ implementation is the most efficient for this benchmark, achieving the **fastest mean execution time** and **lowest minimum time**. This is likely due to the direct compilation to machine code and the lack of additional memory management overhead like garbage collection. The Max time, although 11% slower than Python's, it is in the order of magnitude of 0.1 ms, which may be affected to the the conditions being run a personal PC. This negligible difference at max time to Python may be due to the memory allocation of python bindings being the same order of magnitude.

### Conclusion

While Go is often seen as a performant alternative to Python, our benchmarks reveal that, for OpenCV-based image processing, Python’s mature bindings and efficient C++ backend provide superior performance in terms of both speed and stability. However, Go still is a faster language comparered to Python if Python is not using C++ bindings.

From this quick implementation and set of tests, the hypothesis that simply changing the language would result in faster execution times was not confirmed. Changing the entire codebase to a different language may not be worth the effort due to the increased complexity in implementation and higher development cost. If execution performance is the primary focus, C++ might be a better choice, while Python remains advantageous for rapid development and prototyping due to its ease of use and mature ecosystem.
