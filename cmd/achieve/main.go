package main

import (
	"log"
	"net/http"

	"github.com/topboyasante/achieve/core"
	"github.com/topboyasante/achieve/core/database"
	"github.com/topboyasante/achieve/models"
)

func main() {
	config := core.NewConfig()

	pg, err := database.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	err = database.RunMigrations(pg, &models.Transaction{}, &models.User{}, &models.Wallet{})
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	srv := &http.Server{
		Addr:     ":4000",
	}

	log.Printf("Starting server on port 4000")

	err = srv.ListenAndServe()

	log.Fatal(err)

}