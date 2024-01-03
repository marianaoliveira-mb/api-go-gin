package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marianaoliveira-mb/api-go-gin/controllers"
	"github.com/marianaoliveira-mb/api-go-gin/database"
	"github.com/stretchr/testify/assert"
)

func SetapDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetapDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/Mariana", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai Mariana, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestGetAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetapDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

	fmt.Println(resposta.Body)
}
