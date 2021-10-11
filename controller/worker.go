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

func Reader(c *gin.Context) {
	typeid := c.Param("type")
	items := c.Param("items")
	itemsPerWorker := c.Param("items_per_worker")

	fmt.Println(typeid, items, itemsPerWorker)
}
