# GoArraySortServer

## Overview

**GoArraySortServer** is a Go project that demonstrates sequential and concurrent array sorting using a Go server. The server provides two endpoints, `/process-single` and `/process-concurrent`, allowing users to experience the efficiency and performance differences between sequential and concurrent processing.

## Key Features

1. **Server Setup:**
   - Go server listening on port 8000.
   - Endpoints for sequential (/process-single) and concurrent (/process-concurrent) array processing.

2. **Input Format:**
   - JSON payload with an array structure to be sorted.

3. **Task Implementation:**
   - `/process-single` sorts each sub-array sequentially.
   - `/process-concurrent` sorts each sub-array concurrently using Go's concurrency features.

4. **Response Format:**
   - JSON response with sorted arrays and time taken in nanoseconds.

5. **Performance Measurement:**
   - Measures the time taken to sort all sub-arrays in each endpoint.

6. **Dockerization:**
   - Dockerized for easy deployment and distribution.

## Usage

1. Clone the repository: `git clone git@github.com:bhambhuvikas7376/GoArraySortServer.git`
2. Build the Docker image: `docker build -t goserversortarrays .`
3. Run the Docker container: `docker run -p 8000:8000 goserversortarrays`

## Example `curl` Commands and Expected Output In `json`

### Sequential Processing

**Command:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 1, 2], [6, 4, 5], [9, 7, 8]]}' http://localhost:8000/process-single
```
**Output:**
```json
{"sorted_arrays":[[1,2,3],[4,5,6],[7,8,9]],"time_ns":<time_taken_in_nanoseconds>}
```

### Concurrent Processing

**Command:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 1, 2], [6, 4, 5], [9, 7, 8]]}' http://localhost:8000/process-concurrent
```
**Output:**
```json
{"sorted_arrays":[[1,2,3],[4,5,6],[7,8,9]],"time_ns":<time_taken_in_nanoseconds>}
```

