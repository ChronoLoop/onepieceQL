package categories

import (
	"net/url"

	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func processHref(s string) string {
	fullUrl, err := url.Parse(s)

	if err != nil {
		return s + "?action=render"
	} else {
		return BASE_URL + fullUrl.Path + "?action=render"
	}
}

func processCategoryString(s string) string {
	return utils.RemoveExtraSpaces(utils.RemoveTextWithBrackets(utils.RemoveNewLine(s)))
}

func processFruitEnglishName(s string) string {
	return utils.RemoveAfterFirstFruit(processCategoryString(s))
}
