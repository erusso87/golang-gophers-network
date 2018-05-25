package main

import (
	"log"
	"net/http"
	"fmt"
	"gophers-network/db"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"time"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func gophersHandler(writer http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listGophersHandler(writer, r)
	case "POST":
		createGopherHandler(writer, r)
	}
}


func main() {
	// Loading configuration from ENV
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	http.HandleFunc("/gophers", gophersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}