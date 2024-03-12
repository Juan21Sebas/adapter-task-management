package http

import (
	"net/http"
	"strconv"

	entity "adapter-task-management/internal/core/domain/repository"

	"github.com/gin-gonic/gin"

	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	"adapter-task-management/internal/core/ports"
	//services "adapter-task-management/internal/core/services"
)

type topsecretHandler struct {
	taskService    ports.CommunicationServices
	taskRepository ports.DBRepository
}

func newHandler(service ports.CommunicationServices, repo ports.DBRepository) *topsecretHandler {
	return &topsecretHandler{
		taskService:    service,
		taskRepository: repo,
	}
}

func (o *topsecretHandler) postTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
			return
		}

		entityResponse, err := o.taskService.CreateTask(c, &task)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		entityResponse.Result.Details = []entity.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}
