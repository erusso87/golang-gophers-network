package graphql

import (
	"github.com/graphql-go/graphql"
	"gophers-network/gophers"
	"gophers-network/images"
)

func createRootQuery(
	gopherRepository gophers.GopherRepository,
	imagesRepository images.ImageRepository,
) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"gophers": &graphql.Field{
				Type:        graphql.NewList(gopherType),
				Description: "Get all the gophers",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return gopherRepository.ListGophers()
				},
			},
		},
	})
}
