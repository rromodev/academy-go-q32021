package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkerServiceReader interface {
	Reader(typeId string, items int, items_per_workers int, workers int) (string, error)
}

type WorkerController struct {
	workerService WorkerServiceReader
}

func NewWorkerController(workerService WorkerServiceReader) WorkerController {
	return WorkerController{workerService}
}

func (wc WorkerController) Reader(c *gin.Context) {
	typeid := c.DefaultQuery("type", "odd")
	items, _ := strconv.Atoi(c.DefaultQuery("items", "5"))
	itemsPerWorker, _ := strconv.Atoi(c.DefaultQuery("items_per_workers", "5"))
	workers, _ := strconv.Atoi(c.DefaultQuery("workers", "10"))

	ss, _ := wc.workerService.Reader(typeid, items, itemsPerWorker, workers)
	fmt.Println("controller: ", typeid, items, itemsPerWorker, ss)

	c.JSON(http.StatusOK, ss)
}
