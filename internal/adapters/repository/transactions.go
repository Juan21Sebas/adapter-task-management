package repository

import (
	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"

	schemaget "adapter-task-management/internal/core/domain/repository/schema/getTask"
	schemaUpdate "adapter-task-management/internal/core/domain/repository/schema/updateTask"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *BDRepository) CreateTask(ctx *gin.Context, request *model.Task) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := uuid.NewString()

	query := "INSERT INTO tasks (id, title, description, status) VALUES (?, ?, ?, ?)"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, request.Title, request.Description, request.Status)
	if err != nil {
		return "", err
	}

	return "Su ID es " + id, nil
}

func (p *BDRepository) SelectTask(ctx *gin.Context, request *modelget.Task) (*schemaget.TaskResponse, error) {

	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT title, description, status FROM tasks WHERE id = ?"
	row := p.db.QueryRow(query, request.Id)
	var title, description, status string
	err := row.Scan(&title, &description, &status)
	if err != nil {
		return nil, err
	}
	response := &schemaget.TaskResponse{
		Title:       title,
		Description: description,
		Status:      status,
	}

	return response, nil
}

func (p *BDRepository) UpdateTask(ctx *gin.Context, request *modelupdate.Task) (*schemaUpdate.TaskResponse, error) {

	p.mu.Lock()
	defer p.mu.Unlock()

	query := "UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(request.Title, request.Description, request.Status, request.Id)
	if err != nil {
		return nil, err
	}

	updatedQuery := "SELECT title, description, status FROM tasks WHERE id = ?"
	updatedRow := p.db.QueryRow(updatedQuery, request.Id)
	var updatedTitle, updatedDescription, updatedStatus string
	err = updatedRow.Scan(&updatedTitle, &updatedDescription, &updatedStatus)
	if err != nil {
		return nil, err
	}

	response := &schemaUpdate.TaskResponse{
		Title:       updatedTitle,
		Description: updatedDescription,
		Status:      updatedStatus,
	}

	return response, nil
}

func (p *BDRepository) DeleteTask(ctx *gin.Context, request *modeldelete.Task) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	query := "DELETE FROM tasks WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(request.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		return nil
	}

	return nil
}
