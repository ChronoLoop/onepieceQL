package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/ikevinws/onepieceQL/internal/server/models"
	"github.com/ikevinws/onepieceQL/internal/server/types"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func GetDevilFruitQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.DevilFruitType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)
			devilFruit, err := models.FindDevilFruitByID(id)
			if err != nil {
				return nil, err
			}

			err = models.SignDevilFruitAvatar(&devilFruit)
			if err != nil {
				return nil, err
			}

			return devilFruit, nil
		},
	}
}

func GetDevilFruitsQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "devilFruits",
			Fields: graphql.Fields{
				"info":    &graphql.Field{Type: types.PaginationInfoType},
				"results": &graphql.Field{Type: graphql.NewList(types.DevilFruitType)},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"filter": &graphql.ArgumentConfig{
				Type: types.DevilFruitFilterInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			findDevilFruitsArgs := populateFindDevilFruitsArgs(p.Args)
			page, ok := p.Args["page"].(int)
			if !ok {
				page = 1
			}
			devilFruits, err := models.FindDevilFruits(&findDevilFruitsArgs, page)
			if err != nil {
				return nil, err
			}

			err = models.SignDevilFruitAvatars(&devilFruits)
			if err != nil {
				return nil, err
			}

			count, err := models.CountDevilFruits(&findDevilFruitsArgs)
			if err != nil {
				return nil, err
			}
			pageInfo := utils.CreatePaginationInfo(count, page)

			type GetDevilFruitsQueryResult struct {
				Info    utils.PaginationInfo `json:"info"`
				Results interface{}          `json:"results"`
			}

			return GetDevilFruitsQueryResult{
				Info:    pageInfo,
				Results: devilFruits,
			}, nil
		},
	}
}

func populateFindDevilFruitsArgs(args map[string]interface{}) models.FindDevilFruitsArgs {
	var input models.FindDevilFruitsArgs

	filter, ok := args["filter"].(map[string]interface{})
	if !ok {
		return input
	}

	if name, ok := filter["name"].(string); ok {
		input.EnglishName = name
		input.JapaneseName = name
	}

	if previousOwner, ok := filter["previousOwner"].(string); ok {
		input.PreviousOwner = previousOwner
	}

	if currentOwner, ok := filter["currentOwner"].(string); ok {
		input.CurrentOwner = currentOwner
	}

	if fruitType, ok := filter["type"].(string); ok {
		input.Type = fruitType
	}

	return input
}
