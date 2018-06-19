package graphql

import (
	"github.com/graphql-go/graphql"
	"time"
	"gophers-network/gophers"
	"gophers-network/utils"
	"gophers-network/images"
)

func createRootMutation(
	gopherRepository gophers.GopherRepository,
	imagesRepository images.ImageRepository,
) *graphql.Object {
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
			"createGopherImage": &graphql.Field{
				Type:        imageType,
				Description: "Create new image for Gopher",
				Args: graphql.FieldConfigArgument{
					"gopher": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					gopherId, _ := params.Args["gopher"].(string)
					content, _ := params.Args["content"].(string)

					image := images.Image{
						ID:        utils.CreateULID().String(),
						Content:   content,
						CreatedAt: time.Now().UTC(),
					}

					err := imagesRepository.InsertImage(image)
					if err != nil {
						panic(err)
					}

					gopher := gopherRepository.GetGopher(gopherId)
					gopher.Image = image.ID
					gopherRepository.UpdateGopher(*gopher)

					return image, err
				},
			},
		},
	})
}
