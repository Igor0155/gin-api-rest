package routes

import (
	"github.com/Igor0155/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibiTodosAluno)
	r.GET(":nome", controllers.HelloWorldName )
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.Run()
}