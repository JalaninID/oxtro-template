package main

import (
	"template/config"

	"github.com/JalaninID/workspaces/clients"
	"github.com/JalaninID/workspaces/users"
)

func main() {
	db := config.ConnectToPostgresDB()
	defer db.Close()
	clients.NewClient(db)
	users.NewUser()
}
