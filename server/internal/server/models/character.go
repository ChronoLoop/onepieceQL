package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ikevinws/onepieceQL/internal/server/awsclient"
	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

type Character struct {
	ID             int          `db:"id" json:"id"`
	JapaneseName   string       `db:"japanese_name" csv:"japanese_name" json:"japaneseName"`
	EnglishName    string       `db:"english_name" csv:"english_name" json:"englishName"`
	Debut          string       `db:"debut" csv:"debut" json:"debut"`
	Affiliations   string       `db:"affiliations" csv:"affiliations" json:"affiliations"`
	Origin         string       `db:"origin" csv:"origin" json:"origin"`
	Age            string       `db:"age" csv:"age" json:"age"`
	Birthday       string       `db:"birthday" csv:"birthday" json:"birthday"`
	BloodType      string       `db:"blood_type" csv:"blood_type" json:"bloodType"`
	Bounty         string       `db:"bounty" csv:"bounty" json:"bounty"`
	Description    string       `db:"description" csv:"description" json:"description"`
	DevilFruitName string       `db:"devil_fruit_name" csv:"devil_fruit_name" json:"devilFruitName"`
	AvatarSrc      string       `db:"avatar_src" csv:"avatar_src" json:"avatarSrc"`
	CreatedAt      time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

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

const findCharacterByIDQuery = `
	SELECT *
	FROM characters
	WHERE id = $1 LIMIT 1
`

func FindCharacterByID(id int) (Character, error) {
	character := Character{}
	if err := db.Connection.Get(&character, findCharacterByIDQuery, id); err != nil {
		return character, errors.New("could not find character by id")
	}
	return character, nil
}

const countCharactersQuery = `
	SELECT COUNT(*)
	FROM characters
	WHERE  
	(($1 = '' OR english_name ILIKE '%' || $1 || '%') OR ($2 = '' OR japanese_name ILIKE '%' || $2 || '%'))
	AND ($3 = '' OR affiliations ILIKE '%' || $3 || '%')
	AND ($4 = '' OR origin ILIKE '%' || $4 || '%')
	AND ($5 = '' OR blood_type = $5)
	AND ($6 = '' OR birthday ILIKE '%' || $6 || '%')
	AND ($7 = '' OR devil_fruit_name ILIKE '%' || $7 || '%')
`

func CountCharacters(args *FindCharactersArgs) (int, error) {
	var count int
	err := db.Connection.Get(&count, countCharactersQuery, args.EnglishName, args.JapaneseName, args.Affiliations, args.Origin, args.BloodType, args.Birthday, args.DevilFruitName)
	if err != nil {
		return 0, errors.New("could not count characters")
	}

	return count, nil
}

const findCharactersQuery = `
	SELECT *
	FROM characters
	WHERE  
	(($1 = '' OR english_name ILIKE '%' || $1 || '%') OR ($2 = '' OR japanese_name ILIKE '%' || $2 || '%'))
	AND ($3 = '' OR affiliations ILIKE '%' || $3 || '%')
	AND ($4 = '' OR origin ILIKE '%' || $4 || '%')
	AND ($5 = '' OR blood_type = $5)
	AND ($6 = '' OR birthday ILIKE '%' || $6 || '%')
	AND ($7 = '' OR devil_fruit_name ILIKE '%' || $7 || '%')
	ORDER BY english_name ASC
`

type FindCharactersArgs struct {
	JapaneseName   string `db:"japanese_name"`
	EnglishName    string `db:"english_name"`
	Affiliations   string `db:"affiliations"`
	Origin         string `db:"origin"`
	Birthday       string `db:"birthday"`
	BloodType      string `db:"blood_type"`
	DevilFruitName string `db:"devil_fruit_name"`
}

func FindCharacters(args *FindCharactersArgs, page int) ([]Character, error) {
	characters := []Character{}
	err := db.Connection.Select(
		&characters,
		findCharactersQuery+fmt.Sprintf(" LIMIT %d OFFSET %d", utils.PAGE_LIMIT, (page-1)*utils.PAGE_LIMIT),
		args.EnglishName, args.JapaneseName, args.Affiliations, args.Origin, args.BloodType, args.Birthday, args.DevilFruitName,
	)
	if err != nil {
		return characters, errors.New("could not find characters")
	}
	return characters, nil
}

func SignCharacterAvatar(character *Character) error {
	signedURL, error := awsclient.CloudFrontSigner.Sign(
		awsclient.CLOUDFRONT_DOMAIN_NAME+character.AvatarSrc,
		time.Now().Add(AVATAR_SRC_EXPIRATION_TIME),
	)
	character.AvatarSrc = signedURL
	return error
}

func SignCharacterAvatars(characters *[]Character) error {
	for i := range *characters {
		error := SignCharacterAvatar(&(*characters)[i])
		if error != nil {
			return error
		}
	}
	return nil
}
