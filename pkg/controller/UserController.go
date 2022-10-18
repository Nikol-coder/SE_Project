package Controller

import (
	"SE_Project/pkg/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginForm model.LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		return
	}

}
