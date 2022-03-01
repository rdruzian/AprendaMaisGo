package Autentication

import (
	auth "../Auth"
	"github.com/gin-gonic/gin"
)

func Autentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Autorization")
		if header == "" {
			c.AbortWithStatus(401)
		}
		token := header[len(Bearer_schema):]

		if !auth.NewJWTService().ValidaToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
