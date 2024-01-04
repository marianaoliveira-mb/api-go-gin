package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marianaoliveira-mb/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibeIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run()
}
