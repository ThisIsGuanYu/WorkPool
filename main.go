package main

import (
	"WorkPool/refrigerator"
	"fmt"
	"time"
)

const (
	numBeef    = 10
	numPork    = 7
	numChicken = 5
	TSBeefWasted   = 1 * time.Second
	TSPorkWasted = 2 * time.Second
	TSChickenWasted = 3 * time.Second	
	done = 0 
)

var team = [...]string{"A","B","C","D","E"}

func worker(id string, jobs <-chan refrigerator.Meat, isDone chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %s , Get the %s , At %s\n", id, j.Name, time.Now().Format(time.RFC3339))
		// waiting  opertation 
		time.Sleep(j.OperateTime)
		fmt.Printf("Worker %s , %s is done , At %s\n", id, j.Name, time.Now().Format(time.RFC3339))
		isDone <- done
	}
}

func main() {

	// new a refrigerator
	newRefrigerator := refrigerator.NewRefrigerator()

	// Add meat
	newRefrigerator.Add("Beef",TSBeefWasted,numBeef)
	newRefrigerator.Add("Pork",TSPorkWasted,numPork)
	newRefrigerator.Add("Chicken",TSChickenWasted,numChicken)
	
	numJobs := numBeef + numPork + numChicken

	// jobs are all the meat need to handle
	jobs := make(chan refrigerator.Meat, numJobs)

	// if the meat has been cooked , pass the task done  
	isDone := make(chan int, numJobs)

	// workers get ready by starting goroutine 
	for _ , name := range team {
		go worker(name, jobs, isDone)
	}

	// parpare the meat, move them to the job channel as queue
	for j := 1; j <= numJobs; j++ {
		jobs <- newRefrigerator.Get()
	}
	close(jobs) // the channel could not receive data anymore 

	// check and wait all jobs are done
	for a := 1; a <= numJobs; a++ {
		<-isDone
	}
}
