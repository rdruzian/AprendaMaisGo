package apiUser

import (
	"../../User"
	"../../database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cadastro(c *gin.Context) {
	db := database.GetDatabase()

	var user User.Usuario

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON: " + err.Error(),
		})
		return
	}

	result := db.Create(&user).Error
	if result != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Não foi possível salvar o cadastrar o usuário: %v", result),
		})
	}

	c.JSON(200, gin.H{
		"menssagem": "Usuário cadastro com sucesso!",
	})
}

