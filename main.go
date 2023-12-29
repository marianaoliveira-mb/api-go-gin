package main

import (
	"github.com/marianaoliveira-mb/api-go-gin/database"
	"github.com/marianaoliveira-mb/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
