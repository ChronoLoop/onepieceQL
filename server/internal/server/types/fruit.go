package types

import (
	"github.com/graphql-go/graphql"
)

var DevilFruitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "devilFruit",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"japaneseName":  &graphql.Field{Type: graphql.String},
		"englishName":   &graphql.Field{Type: graphql.String},
		"meaning":       &graphql.Field{Type: graphql.String},
		"usageDebut":    &graphql.Field{Type: graphql.String},
		"type":          &graphql.Field{Type: graphql.String},
		"previousOwner": &graphql.Field{Type: graphql.String},
		"currentOwner":  &graphql.Field{Type: graphql.String},
		"description":   &graphql.Field{Type: graphql.String},
		"avatarSrc":     &graphql.Field{Type: graphql.String},
	},
})

var DevilFruitFilterInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "devilFruitFilter",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"previousOwner": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"currentOwner": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
