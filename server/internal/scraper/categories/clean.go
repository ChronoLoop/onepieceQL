package categories

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func CleanUpFruits() {
	csvPath := "./data/fruits.csv"

	devilFruitsCSV, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer devilFruitsCSV.Close()

	devilFruits := []*csvmodels.DevilFruitCSV{}
	if err := gocsv.UnmarshalFile(devilFruitsCSV, &devilFruits); err != nil {
		panic(err)
	}

	for _, devilFruit := range devilFruits {
		devilFruit.JapaneseName = processCategoryString(devilFruit.JapaneseName)
		devilFruit.EnglishName = processFruitEnglishName(devilFruit.EnglishName)
		devilFruit.Meaning = processCategoryString(devilFruit.Meaning)
		devilFruit.UsageDebut = processCategoryString(devilFruit.UsageDebut)
		devilFruit.Type = processCategoryString(devilFruit.Type)
		devilFruit.PreviousOwner = processCategoryString(devilFruit.PreviousOwner)
		devilFruit.CurrentOwner = processCategoryString(devilFruit.CurrentOwner)
		devilFruit.Description = processCategoryString(devilFruit.Description)
	}

	utils.SaveCSVFile(&devilFruits, csvPath)
}

func CleanUpCharacters() {
	csvPath := "./data/characters.csv"

	charactersCSV, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer charactersCSV.Close()

	characters := []*csvmodels.CharacterCSV{}
	if err := gocsv.UnmarshalFile(charactersCSV, &characters); err != nil {
		panic(err)
	}

	for _, character := range characters {
		character.JapaneseName = processCategoryString(character.JapaneseName)
		character.EnglishName = processCategoryString(character.EnglishName)
		character.Debut = processCategoryString(character.Debut)
		character.Affiliations = processCategoryString(character.Affiliations)
		character.Origin = processCategoryString(character.Origin)
		character.Age = processCategoryString(character.Age)
		character.Birthday = processCategoryString(character.Birthday)
		character.BloodType = processCategoryString(character.BloodType)
		character.Bounty = processCategoryString(character.Bounty)
		character.Description = processCategoryString(character.Description)
		character.DevilFruitName = processCategoryString(character.DevilFruitName)
	}

	utils.SaveCSVFile(&characters, csvPath)
}
