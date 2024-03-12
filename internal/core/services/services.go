package service

import (
	"adapter-task-management/internal/core/ports"
	"net/http"
	"strconv"

	entity "adapter-task-management/internal/core/domain/repository"
	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"

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

func (r *Repository) SelectTask(ctx *gin.Context, request *modelget.Task) (*entity.Response, error) {

	resp, err := r.repo.SelectTask(ctx, request)
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
			Source: "Select Task",
		},
	}, nil

}

func (r *Repository) UpdateTask(ctx *gin.Context, request *modelupdate.Task) (*entity.Response, error) {

	resp, err := r.repo.UpdateTask(ctx, request)
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
			Source: "Update Task",
		},
	}, nil

}

func (r *Repository) DeleteTask(ctx *gin.Context, request *modeldelete.Task) (*entity.Response, error) {

	resp, err := r.repo.DeleteTask(ctx, request)
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
			Source: "Delete Task",
		},
	}, nil

}
