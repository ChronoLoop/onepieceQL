package scraper

import (
	"github.com/ikevinws/onepieceQL/internal/scraper/categories"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func processCategoryString(s string) string {
	return utils.RemoveTextWithBrackets(utils.RemoveNewLine(s))
}

func ScrapeCategories() {
	categories.ScrapeFruits()
	categories.ScrapeCharacters()
	// categories.CleanUpFruits()
	// categories.CleanUpCharacters()
}
