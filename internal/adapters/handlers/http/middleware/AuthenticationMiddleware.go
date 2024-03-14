package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware es un middleware para validar tokens JWT
type JWTMiddleware struct {
	secretKey string
}

// NewJWTMiddleware crea una nueva instancia de JWTMiddleware con la clave secreta proporcionada
func NewJWTMiddleware(secretKey string) *JWTMiddleware {
	return &JWTMiddleware{
		secretKey: secretKey,
	}
}

// MiddlewareFunc es el middleware func para Gin
func (jwtMiddleware *JWTMiddleware) MiddlewareFunc(c *gin.Context) {
	// Obtener el token del encabezado Authorization
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización no proporcionado"})
		return
	}

	// El token debe tener el formato "Bearer token"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		return
	}

	// Validar el token JWT
	tokenString := tokenParts[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtMiddleware.secretKey), nil
	})
	if err != nil {
		log.Println("Error al analizar el token:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
		return
	}

	// Token válido, continuar con el siguiente manejador
	c.Next()
}
