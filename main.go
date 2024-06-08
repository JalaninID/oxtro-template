package main

import (
	c "template/config"

	"github.com/JalaninID/workspaces/clients"
	"github.com/JalaninID/workspaces/users"
)

func main() {
	db := c.ConnectToPostgresDB()
	defer db.Close()
	clients.NewClient(db)
	users.NewUser()
}
