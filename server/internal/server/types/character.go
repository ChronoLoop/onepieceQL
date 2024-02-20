package types

import (
	"github.com/graphql-go/graphql"
)

var CharacterFilterInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "characterFilter",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"affiliations": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"origin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"birthday": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bloodType": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"devilFruitName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var CharacterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "character",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"englishName":    &graphql.Field{Type: graphql.String},
		"japaneseName":   &graphql.Field{Type: graphql.String},
		"origin":         &graphql.Field{Type: graphql.String},
		"debut":          &graphql.Field{Type: graphql.String},
		"affiliations":   &graphql.Field{Type: graphql.String},
		"age":            &graphql.Field{Type: graphql.String},
		"birthday":       &graphql.Field{Type: graphql.String},
		"bloodType":      &graphql.Field{Type: graphql.String},
		"bounty":         &graphql.Field{Type: graphql.String},
		"description":    &graphql.Field{Type: graphql.String},
		"devilFruitName": &graphql.Field{Type: graphql.String},
		"avatarSrc":      &graphql.Field{Type: graphql.String},
	},
})
