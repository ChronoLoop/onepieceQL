package models

import (
	"errors"

	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
)

func DevilFruitTableExists() (bool, error) {
	return tableExists(db.Connection, "devil_fruits")
}

func CreateDevilFruit(devilFruit *csvmodels.DevilFruitCSV) error {
	const query = `
		INSERT INTO devil_fruits (
			japanese_name,
			english_name,
			meaning,
			usage_debut,
			type,
			description,
			current_owner,
			avatar_src 
		) VALUES (	
			:japanese_name, 
			:english_name, 
			:meaning,
			:usage_debut,
			:type,
			:description,
			:current_owner,
			:avatar_src
		)
	`
	if _, err := db.Connection.NamedExec(query, devilFruit); err != nil {
		return errors.New("devil fruit could not be created")
	}
	return nil
}

const checkDevilFruitExistsQuery = `
	SELECT EXISTS (SELECT 1 FROM devil_fruits WHERE english_name = $1 AND avatar_src = $2)
`

func CheckDevilFruitExists(devilFruit *csvmodels.DevilFruitCSV) (bool, error) {
	var exists bool
	if err := db.Connection.Get(&exists, checkDevilFruitExistsQuery, devilFruit.EnglishName, devilFruit.AvatarSrc); err != nil {
		return exists, err
	}
	return exists, nil
}
