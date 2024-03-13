package ports

import (
	entity "adapter-task-management/internal/core/domain/repository"
	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"
	schemaget "adapter-task-management/internal/core/domain/repository/schema/getTask"
	schemaUpdate "adapter-task-management/internal/core/domain/repository/schema/updateTask"

	"github.com/gin-gonic/gin"
)

type CommunicationServices interface {
	CreateTask(ctx *gin.Context, request *model.Task) (*entity.Response, error)
	SelectTask(ctx *gin.Context, request *modelget.Task) (*entity.Response, error)
	UpdateTask(ctx *gin.Context, request *modelupdate.Task) (*entity.Response, error)
	DeleteTask(ctx *gin.Context, request *modeldelete.Task) (*entity.Response, error)
}

type DBRepository interface {
	CreateTask(ctx *gin.Context, request *model.Task) (string, error)
	SelectTask(ctx *gin.Context, request *modelget.Task) (*schemaget.TaskResponse, error)
	UpdateTask(ctx *gin.Context, request *modelupdate.Task) (*schemaUpdate.TaskResponse, error)
	DeleteTask(ctx *gin.Context, request *modeldelete.Task) error
}
