package main

import (
	"log"
	"social/internal/env"
	"social/internal/storage"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	storage := storage.NewPostgresStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
