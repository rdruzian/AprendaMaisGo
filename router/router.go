package router

import (
	Usario "../User/api"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")//, authentication.Auth()) // Tudo feito :)
		{
			user.POST("/cadastro", Usario.Cadastro)
		}
	}

	return router
}
