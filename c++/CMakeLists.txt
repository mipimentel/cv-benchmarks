cmake_minimum_required(VERSION 3.10)
project(CoinSegmentationBenchmark)

find_package(OpenCV REQUIRED)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(benchmark_coins src/benchmark_coins.cpp)
target_link_libraries(benchmark_coins ${OpenCV_LIBS})
