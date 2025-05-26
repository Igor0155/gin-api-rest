package controllers

import (
	"net/http"

	"github.com/Igor0155/api-go-gin/database"
	"github.com/Igor0155/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibiTodosAluno(c *gin.Context){
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func HelloWorldName( c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Hello World, nome passado: " : nome,
	})

}

func CriarNovoAluno(c *gin.Context){

	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}
