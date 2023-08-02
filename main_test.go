package main

import (
	"WorkPool/refrigerator"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	// Set up the refrigerator
	newRefrigerator := refrigerator.NewRefrigerator()
	newRefrigerator.Add("Beef", TSBeefWasted, numBeef)
	newRefrigerator.Add("Pork", TSPorkWasted, numPork)
	newRefrigerator.Add("Chicken", TSChickenWasted, numChicken)

	numJobs := numBeef + numPork + numChicken

	// Create channels
	jobs := make(chan refrigerator.Meat, numJobs)
	isDone := make(chan int, numJobs)

	// Start workers
	for _, name := range team {
		go worker(name, jobs, isDone)
	}

	// Prepare jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- newRefrigerator.Get()
	}
	close(jobs)

	// Wait for all jobs to be done
	for a := 1; a <= numJobs; a++ {
		<-isDone
	}
}

func TestMain(m *testing.M) {
	// Run the tests
	m.Run()
}
