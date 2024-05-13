package main

import (
	"log"
	"net/http"

	"github.com/topboyasante/achieve/data"
	"github.com/topboyasante/achieve/config"
)

func main() {
	config := config.NewConfig()

	pg, err := data.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	err = data.RunMigrations(pg, &data.Transaction{}, &data.User{}, &data.Wallet{})
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