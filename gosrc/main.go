package main

import (
	"ipc/db"
	"ipc/server"
)

func main() { 

db.MigrateDatabase()
server.Server()


}