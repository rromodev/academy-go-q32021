package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type WorkerServiceReader interface {
	Reader(typeId string, items int, items_per_workers int) (string, error)
}

type WorkerController struct {
	workerService WorkerServiceReader
}

func NewWorkerController(workerService WorkerServiceReader) WorkerController {
	return WorkerController{workerService}
}

func (wc WorkerController) Reader(c *gin.Context) {
	typeid := c.DefaultQuery("type", "odd")
	items := 5          // c.DefaultQuery("items", "5")
	itemsPerWorker := 5 // c.DefaultQuery("items_per_worker", "5")

	wc.workerService.Reader(typeid, items, itemsPerWorker)
	fmt.Println(typeid, items, itemsPerWorker)
}
