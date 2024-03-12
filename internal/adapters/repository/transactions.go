package repository

import (
	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"

	schemaget "adapter-task-management/internal/core/domain/repository/schema/getTask"
	schemaUpdate "adapter-task-management/internal/core/domain/repository/schema/updateTask"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (p BDRepository) CreateTask(ctx *gin.Context, request *model.Task) (string, error) {

	query := "INSERT INTO tasks (title, description, status) VALUES (?, ?, ?) RETURNING id"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(request.Title, request.Description, request.Status)
	if err != nil {
		return "", err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return "Registro creado en base de datos y su ID es " + strconv.FormatInt(newID, 10), nil
}

func (p BDRepository) SelectTask(ctx *gin.Context, request *modelget.Task) (*schemaget.TaskResponse, error) {

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

func (p BDRepository) UpdateTask(ctx *gin.Context, request *modelupdate.Task) (*schemaUpdate.TaskResponse, error) {

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

func (p BDRepository) DeleteTask(ctx *gin.Context, request *modeldelete.Task) (string, error) {

	query := "DELETE FROM tasks WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(request.Id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected > 0 {
		return "Registro Eliminado", nil
	}

	return "No se encontró el registro", nil
}
