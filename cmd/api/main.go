package main

import (
	"log"
	"social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
