package main

import (
	"log"
	"fmt"
	"net/http"
	"gophers-network/graphql"
	"github.com/kelseyhightower/envconfig"
	"gophers-network/gophers"
	"github.com/mnmtanish/go-graphiql"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func configFromEnv() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}

func initGopherRepository(config Config) (gopherRepository gophers.GopherRepository) {
	addr := fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
	)

	gopherRepository, err := gophers.CreateRepository(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	return gopherRepository
}

func main() {
	var config = configFromEnv()
	var gopherRepository = initGopherRepository(config)

	// TODO: Repository with infrastructure logic?
	defer gopherRepository.Close()

	http.HandleFunc("/graphql", graphql.CreateRequestSolver(gopherRepository))
	http.HandleFunc("/", graphiql.ServeGraphiQL)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
