package service

import (
	"adapter-task-management/internal/core/ports"
	"net/http"
	"strconv"

	entity "adapter-task-management/internal/core/domain/repository"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"

	"github.com/gin-gonic/gin"
)

type Repository struct {
	repo ports.DBRepository
}

func NewService(repo ports.DBRepository) *Repository {
	return &Repository{
		repo: repo,
	}
}

func (r *Repository) CreateTask(ctx *gin.Context, request *model.Task) (*entity.Response, error) {

	resp, err := r.repo.CreateTask(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Data: resp,
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "",
				},
			},
			Source: "Create Task",
		},
	}, nil

}
