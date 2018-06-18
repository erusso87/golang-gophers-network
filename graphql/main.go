package graphql

import (
	"github.com/graphql-go/graphql"
	"net/http"
	"encoding/json"
	"gophers-network/gophers"
)

type requestData struct {
	Query string
	Variables string
}

func createSchema(gopherRepository gophers.GopherRepository) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    createRootQuery(gopherRepository),
		Mutation: createRootMutation(gopherRepository),
	})
}

func CreateRequestSolver(gopherRepository gophers.GopherRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)

		var requestData requestData
		err := decoder.Decode(&requestData)
		if err != nil {
			panic(err)
		}

		schema, _ := createSchema(gopherRepository)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: requestData.Query,
		})

		json.NewEncoder(writer).Encode(result)
	}
}