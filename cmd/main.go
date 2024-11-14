package main

import (
	"log"

	"github.com/carlosclavijo/loginsolid/cmd/api"
	"github.com/carlosclavijo/loginsolid/internal/config"
	"github.com/carlosclavijo/loginsolid/internal/database"
	"github.com/carlosclavijo/loginsolid/internal/handlers"
)

const portNumber = ":8080"

func main() {
	env := &config.Env{
		Host:     "localhost",
		Port:     "5432",
		Name:     "loginsolid",
		User:     "postgres",
		Password: "abc12345",
	}

	log.Println("Connecting to database...")
	db, err := database.ConnectSQL("host=" + env.Host +
		" port=" + env.Port +
		" dbname=" + env.Name +
		" user=" + env.User +
		" password=" + env.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	server := api.NewApiServer(portNumber, db.SQL)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
