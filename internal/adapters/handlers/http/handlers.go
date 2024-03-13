package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"
	"adapter-task-management/internal/core/ports"
)

type taskManagementHandler struct {
	taskService    ports.CommunicationServices
	taskRepository ports.DBRepository
}

func newHandler(service ports.CommunicationServices, repo ports.DBRepository) *taskManagementHandler {
	return &taskManagementHandler{
		taskService:    service,
		taskRepository: repo,
	}
}

func (o *taskManagementHandler) postTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
			return
		}
		entityResponse, err := o.taskService.CreateTask(c, &task)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *taskManagementHandler) getTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task modelget.Task
		task.Id = c.Param("id")
		if err := c.ShouldBindQuery(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.taskService.SelectTask(c, &task)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *taskManagementHandler) putTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task modelupdate.Task
		task.Id = c.Param("id")
		if err := c.ShouldBindQuery(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.taskService.UpdateTask(c, &task)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

func (o *taskManagementHandler) deleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task modeldelete.Task
		task.Id = c.Param("id")
		if err := c.ShouldBindQuery(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.taskService.DeleteTask(c, &task)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}
