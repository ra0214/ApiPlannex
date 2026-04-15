package infraestructure

import (
	"Plannex/src/users/domain"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, userRepo domain.IUser) {
	ps := NewMySQL()
	SetupRouter(ps, r, userRepo)
}
