package service

import (
	"errors"
	"net/http/httptest"
	"testing"

	entity "adapter-task-management/internal/core/domain/repository"
	db "adapter-task-management/internal/core/ports"

	modelget "adapter-task-management/internal/core/domain/repository/model/getTask"
	model "adapter-task-management/internal/core/domain/repository/model/postTask"
	mockRepository "adapter-task-management/internal/core/ports/mocks"

	modeldelete "adapter-task-management/internal/core/domain/repository/model/deleteTask"
	modelupdate "adapter-task-management/internal/core/domain/repository/model/updateTask"
	getTask "adapter-task-management/internal/core/domain/repository/schema/getTask"
	updateTask "adapter-task-management/internal/core/domain/repository/schema/updateTask"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	var repository db.DBRepository
	service := NewService(repository)
	assert.NotNil(t, service, "Servicio no debe ser nil")
}

func TestCreateTask(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("CreateTask", mock.Anything, mock.Anything).Return("", nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := model.Task{
		Title:       "Tittle",
		Description: "Description",
		Status:      "Status",
	}

	expectedResp := entity.Response{
		Data: "",
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Creado"},
			},
			Source: "Create Task",
		},
	}

	response, err := svc.CreateTask(c, &req)
	assert.NoError(t, err)
	assert.Equal(t, &expectedResp, response)
}

func TestCreateTask_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("CreateTask", mock.Anything, mock.Anything).Return("", errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := model.Task{
		Title:       "Tittle",
		Description: "Description",
		Status:      "Status",
	}

	response, err := svc.CreateTask(c, &req)

	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestSelectTask(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	resultMockService1 := &getTask.TaskResponse{}
	mockRepo.On("SelectTask", mock.Anything, mock.Anything).Return(resultMockService1, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modelget.Task{
		Id: "ID",
	}

	expectedResp := entity.Response{
		Data: resultMockService1,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Seleccionado"},
			},
			Source: "Select Task",
		},
	}

	response, err := svc.SelectTask(c, &req)
	assert.NoError(t, err)
	assert.Equal(t, &expectedResp, response)
}

func TestSelectTask_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	resultMockService1 := &getTask.TaskResponse{}
	mockRepo.On("SelectTask", mock.Anything, mock.Anything).Return(resultMockService1, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modelget.Task{
		Id: "ID",
	}

	response, err := svc.SelectTask(c, &req)

	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	resultMockService1 := &updateTask.TaskResponse{}
	mockRepo.On("UpdateTask", mock.Anything, mock.Anything).Return(resultMockService1, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modelupdate.Task{
		Id:          "ID",
		Title:       "Tittle",
		Description: "Description",
		Status:      "Status",
	}

	expectedResp := entity.Response{
		Data: resultMockService1,
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Actualizado"},
			},
			Source: "Update Task",
		},
	}

	response, err := svc.UpdateTask(c, &req)
	assert.NoError(t, err)
	assert.Equal(t, &expectedResp, response)
}

func TestUpdateTask_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	resultMockService1 := &updateTask.TaskResponse{}
	mockRepo.On("UpdateTask", mock.Anything, mock.Anything).Return(resultMockService1, errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modelupdate.Task{
		Id:          "ID",
		Title:       "Tittle",
		Description: "Description",
		Status:      "Status",
	}

	response, err := svc.UpdateTask(c, &req)

	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("DeleteTask", mock.Anything, mock.Anything).Return(nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modeldelete.Task{
		Id: "ID",
	}

	expectedResp := entity.Response{
		Result: entity.Result{
			Details: []entity.Detail{
				{InternalCode: "200", Message: "OK", Detail: "Registro Eliminado"},
			},
			Source: "Delete Task",
		},
	}

	response, err := svc.DeleteTask(c, &req)
	assert.NoError(t, err)
	assert.Equal(t, &expectedResp, response)
}

func TestDeleteTask_ErrorCase(t *testing.T) {
	mockRepo := mockRepository.NewDBRepository(t)
	svc := &Repository{repo: mockRepo}

	mockRepo.On("DeleteTask", mock.Anything, mock.Anything).Return(errors.New("error simulado"))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := modeldelete.Task{
		Id: "ID",
	}

	response, err := svc.DeleteTask(c, &req)

	assert.Error(t, err)
	assert.Nil(t, response)
}
