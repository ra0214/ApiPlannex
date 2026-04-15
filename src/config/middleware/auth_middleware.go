package middleware

import (
	"Plannex/src/users/domain"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida el token de autenticación y extrae el ID del usuario
func AuthMiddleware(userRepo domain.IUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// El token viene en formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		token := parts[1]

		// Buscar el usuario por token
		user, err := userRepo.GetUserByToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token inválido o expirado")})
			c.Abort()
			return
		}

		// Guardar el ID del usuario en el contexto
		c.Set("userID", user.ID)
		c.Set("user", user)

		c.Next()
	}
}
