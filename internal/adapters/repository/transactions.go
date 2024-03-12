package repository

import (
	model "adapter-task-management/internal/core/domain/repository/model/postTask"

	"github.com/gin-gonic/gin"
)

func (p BDRepository) CreateTask(ctx *gin.Context, request *model.Task) (string, error) {

	_, err := p.db.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)",
		request.Title, request.Description, request.Status)
	if err != nil {
		return "", err
	}

	return "Ok", nil
}
