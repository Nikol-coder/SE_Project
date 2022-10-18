package router

import (
	"SE_Project/pkg/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("Login", Controller.Login)
	return r
}
