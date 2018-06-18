package graphql

import (
	"github.com/graphql-go/graphql"
	"time"
	"gophers-network/gophers"
	"gophers-network/utils"
)

func createRootMutation(gopherRepository gophers.GopherRepository) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createGopher": &graphql.Field{
				Type:        gopherType,
				Description: "Create new gopher",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"color": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"weight": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					color, _ := params.Args["color"].(string)
					weight, _ := params.Args["weight"].(float32)

					gopher := gophers.Gopher{
						ID:        utils.CreateULID().String(),
						Name:      name,
						Color:     color,
						Weight:    weight,
						CreatedAt: time.Now().UTC(),
					}

					err := gopherRepository.InsertGopher(gopher)
					return gopher, err
				},
			},
		},
	})
}
