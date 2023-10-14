package models

import (
	"errors"

	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
)

func CharacterTableExists() (bool, error) {
	return tableExists(db.Connection, "characters")
}

const createCharacterQuery = `
	INSERT INTO characters (
		japanese_name,
		english_name,
		debut,
		affiliations,
		origin,
		age,
		birthday,
		blood_type,
		bounty,
		description,
		devil_fruit_name,
		avatar_src
	) VALUES (	
		:japanese_name, 
		:english_name, 
		:debut,
		:affiliations,
		:origin,
		:age,
		:birthday,
		:blood_type,
		:bounty,
		:description,
		:devil_fruit_name,
		:avatar_src
	)
`

func CreateCharacter(character *csvmodels.CharacterCSV) error {
	if _, err := db.Connection.NamedExec(createCharacterQuery, character); err != nil {
		return errors.New("character could not be created")
	}
	return nil
}

const checkCharacterExistsQuery = `
	SELECT EXISTS (SELECT 1 FROM characters WHERE english_name = $1 AND avatar_src = $2)
`

func CheckCharacterExists(character *csvmodels.CharacterCSV) (bool, error) {
	var exists bool
	if err := db.Connection.Get(&exists, checkCharacterExistsQuery, character.EnglishName, character.AvatarSrc); err != nil {
		return exists, err
	}
	return exists, nil
}
