package main

import (
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/storage"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONS", 25),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONS", 25),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenCons,
		cfg.db.maxIdleCons,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("connected to database %s", cfg.db.addr)

	storage := storage.NewPostgresStorage(db)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
