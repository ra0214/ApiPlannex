package infraestructure

import (
	"Plannex/src/users/application"
	"Plannex/src/users/domain"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	useCase *application.LoginUser
	repo    domain.IUser
}

func NewLoginUserController(useCase *application.LoginUser, repo domain.IUser) *LoginUserController {
	return &LoginUserController{useCase: useCase, repo: repo}
}

type LoginRequestBody struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func generateAuthToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (lc *LoginUserController) Execute(c *gin.Context) {
	var body LoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON"})
		return
	}

	user, err := lc.useCase.Execute(body.UserName, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	token, err := generateAuthToken()
	if err == nil {
		_ = lc.repo.UpdateUserAuthToken(user.ID, token)
		user.AuthToken = token
	}

	// No devolver la contraseña en la respuesta
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    user,
	})
}
