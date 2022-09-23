package main

import (
	"github.com/CaioLuColaco/api-upVote-golang/database"
	"github.com/CaioLuColaco/api-upVote-golang/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}