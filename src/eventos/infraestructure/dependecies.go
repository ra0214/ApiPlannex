package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := NewMySQL()
	SetupRouter(ps, r)
}
