package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

type DevilFruit struct {
	ID            int          `db:"id" json:"id"`
	JapaneseName  string       `db:"japanese_name" csv:"japanese_name" json:"japaneseName"`
	EnglishName   string       `db:"english_name" csv:"english_name" json:"englishName"`
	Meaning       string       `db:"meaning" csv:"meaning" json:"meaning"`
	UsageDebut    string       `db:"usage_debut" csv:"usage_debut" json:"usageDebut"`
	Type          string       `db:"type" csv:"type" json:"type"`
	PreviousOwner string       `db:"previous_owner" csv:"previous_owner" json:"previousOwner"`
	CurrentOwner  string       `db:"current_owner" csv:"current_owner" json:"currentOwner"`
	Description   string       `db:"description" csv:"description" json:"description"`
	AvatarSrc     string       `db:"avatar_src" csv:"avatar_src" json:"avatarSrc"`
	CreatedAt     time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

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
			previous_owner,
			avatar_src 
		) VALUES (	
			:japanese_name, 
			:english_name, 
			:meaning,
			:usage_debut,
			:type,
			:description,
			:current_owner,
            :previous_owner,
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

const findDevilFruitByIDQuery = `
	SELECT *
	FROM devil_fruits
	WHERE id = $1 LIMIT 1
`

func FindDevilFruitByID(id int) (DevilFruit, error) {
	devilFruit := DevilFruit{}
	if err := db.Connection.Get(&devilFruit, findDevilFruitByIDQuery, id); err != nil {
		return devilFruit, errors.New("could not find devil fruit by id")
	}
	return devilFruit, nil
}

type FindDevilFruitsArgs struct {
	EnglishName   string `db:"english_name"`
	JapaneseName  string `db:"japanese_name"`
	PreviousOwner string `db:"previous_owner"`
	CurrentOwner  string `db:"current_owner"`
	Type          string `db:"type"`
}

const findDevilFruitsQuery = `
	SELECT *
	FROM devil_fruits
	WHERE  
	(($1 = '' OR english_name ILIKE '%' || $1 || '%') OR ($2 = '' OR japanese_name ILIKE '%' || $2 || '%'))
	AND ($3 = '' OR previous_owner ILIKE '%' || $3 || '%')
	AND ($4 = '' OR current_owner ILIKE '%' || $4 || '%')
	AND ($5 = '' OR type ILIKE '%' || $5 || '%')
	ORDER BY english_name ASC
`

func FindDevilFruits(args *FindDevilFruitsArgs, page int) ([]DevilFruit, error) {
	devilFruits := []DevilFruit{}
	err := db.Connection.Select(
		&devilFruits,
		findDevilFruitsQuery+fmt.Sprintf(" LIMIT %d OFFSET %d", utils.PAGE_LIMIT, (page-1)*utils.PAGE_LIMIT),
		args.EnglishName, args.JapaneseName, args.PreviousOwner, args.CurrentOwner, args.Type,
	)
	if err != nil {
		println(err.Error())
		return devilFruits, errors.New("could not find devil fruits")
	}
	return devilFruits, nil
}

const countDevilFruitsQuery = `
	SELECT COUNT(*)
	FROM devil_fruits
	WHERE  
	(($1 = '' OR english_name ILIKE '%' || $1 || '%') OR ($2 = '' OR japanese_name ILIKE '%' || $2 || '%'))
	AND ($3 = '' OR previous_owner ILIKE '%' || $3 || '%')
	AND ($4 = '' OR current_owner ILIKE '%' || $4 || '%')
	AND ($5 = '' OR type ILIKE '%' || $5 || '%')
`

func CountDevilFruits(args *FindDevilFruitsArgs) (int, error) {
	var count int
	err := db.Connection.Get(&count, countDevilFruitsQuery, args.EnglishName, args.JapaneseName, args.PreviousOwner, args.CurrentOwner, args.Type)
	if err != nil {
		return 0, errors.New("could not count devil fruits")
	}

	return count, nil
}
