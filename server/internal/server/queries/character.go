package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/ikevinws/onepieceQL/internal/server/models"
	"github.com/ikevinws/onepieceQL/internal/server/types"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func GetCharacterQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.CharacterType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)
			character, err := models.FindCharacterByID(id)
			if err != nil {
				return nil, err
			}
			return character, nil
		},
	}
}

func GetCharactersQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "characters",
			Fields: graphql.Fields{
				"info":    &graphql.Field{Type: types.PaginationInfoType},
				"results": &graphql.Field{Type: graphql.NewList(types.CharacterType)},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"filter": &graphql.ArgumentConfig{
				Type: types.CharacterFilterInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			findCharactersArgs := populateFindCharactersArgs(p.Args)
			page, ok := p.Args["page"].(int)
			if !ok {
				page = 1
			}
			characters, err := models.FindCharacters(&findCharactersArgs, page)
			if err != nil {
				return nil, err
			}
			count, err := models.CountCharacters(&findCharactersArgs)
			if err != nil {
				return nil, err
			}
			pageInfo := utils.CreatePaginationInfo(count, page)

			type GetCharactersQueryResult struct {
				Info    utils.PaginationInfo `json:"info"`
				Results []models.Character   `json:"results"`
			}

			return GetCharactersQueryResult{
				Info:    pageInfo,
				Results: characters,
			}, nil
		},
	}
}

func populateFindCharactersArgs(args map[string]interface{}) models.FindCharactersArgs {
	var input models.FindCharactersArgs

	filter, ok := args["filter"].(map[string]interface{})
	if !ok {
		return input
	}

	if name, ok := filter["name"].(string); ok {
		input.EnglishName = name
		input.JapaneseName = name
	}

	if affiliation, ok := filter["affiliation"].(string); ok {
		input.Affiliations = affiliation
	}

	if origin, ok := filter["origin"].(string); ok {
		input.Origin = origin
	}

	if birthday, ok := filter["birthday"].(string); ok {
		input.Birthday = birthday
	}

	if bloodType, ok := filter["bloodType"].(string); ok {
		input.BloodType = bloodType
	}

	if devilFruitName, ok := filter["devilFruitName"].(string); ok {
		input.DevilFruitName = devilFruitName
	}

	return input
}
