package service

import (
	"fmt"
	"sync"
)

type CSVWorkerSetter interface {
	WorkerReader(csvLine chan<- []string, items int, typeNumber string)
}

type WorkerService struct {
	data CSVWorkerSetter
}

func NewWorkerService(data CSVWorkerSetter) WorkerService {
	return WorkerService{data: data}
}

//const workerPoolSize = 50

func (ws WorkerService) Reader(typeId string, items int, items_per_workers int, workerPoolSize int) (string, error) {
	//jobs := make(chan []string, workerPoolSize)
	csvLine := make(chan []string, workerPoolSize)
	results := make(chan string)

	lines := 0
	someMapMutex := sync.RWMutex{}
	m := make(map[int]int)

	go ws.data.WorkerReader(csvLine, items, typeId)

	wg := &sync.WaitGroup{}
	wg.Add(workerPoolSize)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= workerPoolSize; i++ {
		go func(index int) {
			defer wg.Done()

			for job := range csvLine {

				someMapMutex.Lock()
				if m[index]+1 <= items_per_workers {
					lines++
					m[index]++
					results <- fmt.Sprintf("%d Worker %d starting: %s\n", lines, index, job[0])
				}
				someMapMutex.Unlock()
			}

		}(i)
	}

	for r := range results {
		fmt.Println("result", r)
	}

	return fmt.Sprintf("test, lines %d map %v", lines, m), nil
}
