package graphql

import "github.com/graphql-go/graphql"

var gopherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gopher",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"weight": &graphql.Field{
			Type: graphql.Float,
		},
		// TODO: Set as Image (type)
		"image": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var imageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Image",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})