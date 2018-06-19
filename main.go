package main

import (
	"log"
	"fmt"
	"net/http"
	"gophers-network/graphql"
	"github.com/kelseyhightower/envconfig"
	"gophers-network/gophers"
	"github.com/mnmtanish/go-graphiql"
	"gophers-network/images"
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

func initRepositories(config Config) (gr gophers.GopherRepository, ir images.ImageRepository) {
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

	imageRepository, err := images.CreateRepository(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	return gopherRepository, imageRepository
}

func main() {
	var config = configFromEnv()
	var gophersRepo, imagesRepo = initRepositories(config)

	// TODO: Repository with infrastructure logic?
	// TODO: Dependencies container? (For repos)
	defer gophersRepo.Close()
	defer imagesRepo.Close()

	http.HandleFunc("/graphql", graphql.CreateRequestSolver(gophersRepo, imagesRepo))
	http.HandleFunc("/", graphiql.ServeGraphiQL)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
