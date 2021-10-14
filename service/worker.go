package service

import "fmt"

type CSVWorkerSetter interface {
	WorkerReader(ID int, jobs <-chan int, results chan<- []string)
}

type WorkerService struct {
	data CSVWorkerSetter
}

func NewWorkerService(data CSVWorkerSetter) WorkerService {
	return WorkerService{data: data}
}

const workerPoolSize = 4

func (ws WorkerService) Reader(typeId string, items int, items_per_workers int) (string, error) {
	jobs := make(chan int, 10)
	results := make(chan []string, 10)
	fmt.Println("worker func called")
	for x := 1; x <= 3; x++ {
		go ws.data.WorkerReader(x, jobs, results)
	}

	// Give them jobs
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for the results
	for r := 1; r <= 6; r++ {
		fmt.Println("Result received from worker: ", <-results)
	}

	return "", nil
}
