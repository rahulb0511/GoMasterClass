/*
Implement a worker pool with 3 workers to process 15 jobs concurrently.
Each job should simply print its index and sleep for that many seconds.
Ensure that all jobs are completed by the worker pool before the program exits.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Duration(job) * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		wg.Done()
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 15

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}

	close(jobs) // No more jobs will be sent

	wg.Wait() // Wait for all jobs to finish
	fmt.Println("All jobs completed.")
}
