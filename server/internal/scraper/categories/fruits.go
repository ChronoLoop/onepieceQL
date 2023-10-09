package categories

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

var DEVIL_FRUIT_CATEGORY_LINKS = [...]string{
	"https://onepiece.fandom.com/wiki/Category:Paramecia?action=render",
	"https://onepiece.fandom.com/wiki/Category:Logia?action=render",
	"https://onepiece.fandom.com/wiki/Category:Zoan?action=render",
	"https://onepiece.fandom.com/wiki/Category:Mythical_Zoan?action=render",
	"https://onepiece.fandom.com/wiki/Category:Ancient_Zoan?action=render",
}

type DevilFruit struct {
	JapaneseName string
	EnglishName  string
	Meaning      string
	UsageDebut   string
	Type         string
	PreviousUser string
	CurrentUser  string
	Description  string
}

func isDevilFruit(s string) bool {
	//regexp does not support backreference
	pattern := `(?i)^(\w+)\s(\w+)\sno\s+mi`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(s)
}

func getDevilFruitLinksFromPage(pageURL string, wg *sync.WaitGroup, resultChan chan<- []string) {
	defer wg.Done()
	page := rod.New().MustConnect().MustPage(pageURL)
	aElements := page.MustElement("body > div.category-page__members").MustElements("a.category-page__member-link")
	links := []string{}
	for _, a := range aElements {
		text := a.MustText()
		if !isDevilFruit(text) {
			continue
		}
		href := a.MustAttribute("href")
		if href == nil {
			continue
		}
		links = append(links, processHref(*href))
	}
	resultChan <- links
}

func getDevilFruitLinks() []string {
	var wg sync.WaitGroup
	linksChan := make(chan []string, len(DEVIL_FRUIT_CATEGORY_LINKS))
	for _, link := range DEVIL_FRUIT_CATEGORY_LINKS {
		go getDevilFruitLinksFromPage(link, &wg, linksChan)
		wg.Add(1)
	}
	go func() {
		wg.Wait()
		close(linksChan)
	}()
	var allDevilFruitLinks []string
	for links := range linksChan {
		allDevilFruitLinks = append(allDevilFruitLinks, links...)
	}
	allDevilFruitLinks = utils.RemoveDuplicate[string](allDevilFruitLinks)
	return allDevilFruitLinks
}

func getDevilFruitFromPage(browser *rod.Browser, pageLink string, wg *sync.WaitGroup, resultChan chan<- DevilFruit) {
	defer wg.Done()
	page := browser.MustPage(pageLink)
	asideElement := page.MustElement("body > div > aside")
	japaneseName := asideElement.MustElementR("section h3", "/japanese name/i").MustNext().MustText()
	englishName := "N/A"
	rod.Try(func() {
		if englishNameHeader := asideElement.Timeout(2*time.Second).MustElementR("section h3", "/(official english name|romanized name):/i"); englishNameHeader != nil {
			englishName = englishNameHeader.MustNext().MustText()
		}
	})
	meaning := asideElement.MustElementR("section h3", "/meaning:/i").MustNext().MustText()
	usageDebut := asideElement.MustElementR("section h3", "/usage debut:/i").MustNext().MustText()
	previousUser := "N/A"
	rod.Try(func() {
		if previousUserHeader := asideElement.Timeout(2*time.Second).MustElementR("section h3", "/previous (user|owner):/i"); previousUserHeader != nil {
			previousUser = previousUserHeader.MustNext().MustText()
		}
	})
	currentUser := "N/A"
	rod.Try(func() {
		if currentUserHeader := asideElement.Timeout(2*time.Second).MustElementR("section h3", "/current (user|owner):/i"); currentUserHeader != nil {
			currentUser = currentUserHeader.MustNext().MustText()
		}
	})
	fruitType := asideElement.MustElementR("section h3", "/type:/i").MustNext().MustText()
	description := "N/A"
	for i := 0; i < 3; i++ {
		selector := "body > div > p:nth-child(" + strconv.Itoa(i+4) + ")"
		rod.Try(func() {
			if descriptionParagraphElement := page.Timeout(2 * time.Second).MustElement(selector); descriptionParagraphElement != nil {
				description = descriptionParagraphElement.MustText()
			}
		})
		if description != "N/A" && description != "" {
			break
		}
	}

	devilFruit := DevilFruit{
		JapaneseName: utils.RemoveTextWithBrackets(japaneseName),
		EnglishName:  utils.RemoveTextWithBrackets(englishName),
		Meaning:      utils.RemoveTextWithBrackets(meaning),
		UsageDebut:   utils.RemoveTextWithBrackets(usageDebut),
		Type:         utils.RemoveTextWithBrackets(fruitType),

		PreviousUser: utils.RemoveTextWithBrackets(previousUser),
		CurrentUser:  utils.RemoveTextWithBrackets(currentUser),
		Description:  utils.RemoveTextWithBrackets(description),
	}
	resultChan <- devilFruit
}
func getDevilFruitsFromLinks(pageLinks []string) []DevilFruit {
	devilFruits := []DevilFruit{}
	var wg sync.WaitGroup
	devilFruitChan := make(chan DevilFruit, len(pageLinks))
	browser := rod.New().MustConnect()
	for _, link := range pageLinks {
		go getDevilFruitFromPage(browser, link, &wg, devilFruitChan)
		wg.Add(1)
	}
	go func() {
		wg.Wait()
		close(devilFruitChan)
	}()
	for devilFruit := range devilFruitChan {
		devilFruits = append(devilFruits, devilFruit)
	}

	return devilFruits
}

func createDevilFruitCsv(devilFruits []DevilFruit) {
	devilFruitRows := [][]string{
		{"japanese_name", "english_name", "meaning", "usage_debut", "previous_user", "current_user", "description"},
	}
	for _, devilFruit := range devilFruits {
		row := []string{devilFruit.JapaneseName, devilFruit.EnglishName, devilFruit.Meaning, devilFruit.UsageDebut, devilFruit.PreviousUser, devilFruit.CurrentUser, devilFruit.Description}
		devilFruitRows = append(devilFruitRows, row)
	}
	file, err := os.Create("./data/fruits.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(devilFruitRows); err != nil {
		fmt.Println("Error writing rows to file:", err)
	}
}

func LoadFruits() {
	fmt.Println("Scraping devil fruits...")
	defer utils.Timer("Created fruits.csv")()
	devilFruitLinks := getDevilFruitLinks()
	devilFruits := getDevilFruitsFromLinks(devilFruitLinks)
	createDevilFruitCsv(devilFruits)
}
