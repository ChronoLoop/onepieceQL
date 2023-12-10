package types

import (
	"github.com/graphql-go/graphql"
)

var PaginationInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "pagenationInfo",
	Fields: graphql.Fields{
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"pages": &graphql.Field{
			Type: graphql.Int,
		},
		"next": &graphql.Field{
			Type: graphql.Int,
		},
		"prev": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
