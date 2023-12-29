package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marianaoliveira-mb/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.Run()
}
