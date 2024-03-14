package http

import (
	"adapter-task-management/internal/adapters/handlers/http/middleware"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func CreateServer(db *sql.DB) *gin.Engine {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,DELETE,PUT",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))

	jwtMiddleware := middleware.NewJWTMiddleware("s3cr3tK3yF0rJWT!")

	// Usar el middleware JWT directamente con Gin
	server.Use(jwtMiddleware.MiddlewareFunc)

	RegisterRoutes(server, db)

	return server
}

func RunServer(db *sql.DB) {
	server := CreateServer(db)
	server.Run(":8086")
}
