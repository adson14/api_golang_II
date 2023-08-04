package main

import (
	"bitbucket.org/adson14/api_golang_ii.git/database"
	"bitbucket.org/adson14/api_golang_ii.git/server"
)

func main() {

	database.StartDb()
	server := server.NewServer()

	server.Run()
}
