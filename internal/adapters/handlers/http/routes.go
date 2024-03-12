package http

import (
	repository "adapter-task-management/internal/adapters/repository"
	services "adapter-task-management/internal/core/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine, db *sql.DB) {
	// Crea e inicializa el repositorio BDRepository con la conexión a la base de datos
	taskRepository := repository.NewBdRepository(db)

	// Crea e inicializa el servicio con el repositorio
	taskManagementService := services.NewService(taskRepository)

	// Crea el manejador con el servicio y el repositorio
	taskManagementHandler := newHandler(taskManagementService, taskRepository)

	// Registra las rutas
	e.POST("/task/", taskManagementHandler.postTask())
	e.GET("/task/:id", taskManagementHandler.getTask())
	e.PUT("/task/:id", taskManagementHandler.putTask())
	e.DELETE("/task/:id", taskManagementHandler.deleteTask())
}
