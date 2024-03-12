package ports

import (
	entity "adapter-task-management/internal/core/domain/repository"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"

	"github.com/gin-gonic/gin"
)

type CommunicationServices interface {
	CreateTask(ctx *gin.Context, request *model.Task) (*entity.Response, error)
}

type DBRepository interface {
	CreateTask(ctx *gin.Context, request *model.Task) (string, error)
}
