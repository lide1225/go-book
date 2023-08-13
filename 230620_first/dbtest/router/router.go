package router

import (
	"dbt/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
