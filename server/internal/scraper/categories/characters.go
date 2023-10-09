package categories

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

var CHARACTER_LIST_LINKS = [...]string{
	"https://onepiece.fandom.com/wiki/List_of_Canon_Characters?action=render",
}

type Character struct {
	JapaneseName   string
	EnglishName    string
	Debut          string
	Affiliations   string
	Origin         string
	Age            string
	Birthday       string
	BloodType      string
	Bounty         string
	Description    string
	DevilFruitName string
}

func getCharacterLinksFromPage(pageURL string, wg *sync.WaitGroup, resultChan chan<- []string) {
	defer wg.Done()
	page := rod.New().MustConnect().MustPage(pageURL)
	aElements := page.MustElement("body > div > table:nth-child(7) > tbody").MustElements("tr > td:nth-child(2) > a")
	links := []string{}
	for _, a := range aElements {
		href := a.MustAttribute("href")
		if href == nil {
			continue
		}

		links = append(links, processHref(*href))
	}
	resultChan <- links
}

func getCharacterLinks() []string {
	var wg sync.WaitGroup
	linksChan := make(chan []string, len(CHARACTER_LIST_LINKS))
	for _, link := range CHARACTER_LIST_LINKS {
		go getCharacterLinksFromPage(link, &wg, linksChan)
		wg.Add(1)
	}
	go func() {
		wg.Wait()
		close(linksChan)
	}()
	var characterLinks []string
	for links := range linksChan {
		characterLinks = append(characterLinks, links...)
	}
	characterLinks = utils.RemoveDuplicate[string](characterLinks)
	return characterLinks
}

func getCharacterFromPage(browser *rod.Browser, pageLink string, wg *sync.WaitGroup, resultChan chan<- Character) {

	defer wg.Done()
	page := browser.MustPage(pageLink)
	var asideElement *rod.Element = nil
	rod.Try(func() {
		asideElement = page.Timeout(5 * time.Second).MustElement("body > div > aside")
	})
	if asideElement == nil {
		fmt.Printf("could not find aside element: %s", pageLink)
		resultChan <- Character{}
		return
	}
	var characterSection *rod.Element = nil
	rod.Try(func() {
		characterSection = asideElement.Timeout(1 * time.Second).MustElement("section:nth-of-type(1)")
	})
	if characterSection == nil {
		fmt.Printf("could not find character section element: %s", pageLink)
		resultChan <- Character{}
		return
	}

	japaneseName := "N/A"
	englishName := "N/A"
	debut := "N/A"
	affiliations := "N/A"
	origin := "N/A"
	age := "N/A"
	birthday := "N/A"
	bloodtype := "N/A"
	bounty := "N/A"
	description := "N/A"

	rod.Try(func() {
		japaneseName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/japanese name:/i").MustNext().MustText()
	})
	rod.Try(func() {
		englishName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/english name:/i").MustNext().MustText()
	})
	rod.Try(func() {
		debut = characterSection.Timeout(1*time.Second).MustElementR("h3", "/debut:/i").MustNext().MustText()
	})
	rod.Try(func() {
		affiliations = characterSection.Timeout(1*time.Second).MustElementR("h3", "/affiliations:/i").MustNext().MustText()
	})
	rod.Try(func() {
		origin = characterSection.Timeout(1*time.Second).MustElementR("h3", "/origin:/i").MustNext().MustText()
	})
	rod.Try(func() {
		age = characterSection.Timeout(1*time.Second).MustElementR("h3", "/age:/i").MustNext().MustText()
	})
	rod.Try(func() {
		birthday = characterSection.Timeout(1*time.Second).MustElementR("h3", "/birthday:/i").MustNext().MustText()
	})
	rod.Try(func() {
		bloodtype = characterSection.Timeout(1*time.Second).MustElementR("h3", "/blood type:/i").MustNext().MustText()
	})
	rod.Try(func() {
		bounty = characterSection.Timeout(1*time.Second).MustElementR("h3", "/bounty:/i").MustNext().MustText()
	})

	for i := 0; i < 3; i++ {
		selector := "body > div > p:nth-child(" + strconv.Itoa(i+4) + ")"
		rod.Try(func() {
			if descriptionParagraphElement := page.Timeout(1 * time.Second).MustElement(selector); descriptionParagraphElement != nil {
				description = descriptionParagraphElement.MustText()
			}
		})
		if description != "N/A" && description != "" {
			break
		}
	}

	devilFruitName := "N/A"
	rod.Try(func() {
		characterDevilFruitSection := asideElement.Timeout(1 * time.Second).MustElement("section:nth-of-type(2)")
		devilFruitName = characterDevilFruitSection.Timeout(1*time.Second).MustElementR("h3", "/japanese name:/i").MustNext().MustText()
	})

	character := Character{
		JapaneseName:   utils.RemoveTextWithBrackets(japaneseName),
		EnglishName:    utils.RemoveTextWithBrackets(englishName),
		Debut:          utils.RemoveTextWithBrackets(debut),
		Affiliations:   utils.RemoveTextWithBrackets(affiliations),
		Origin:         utils.RemoveTextWithBrackets(origin),
		Age:            utils.RemoveTextWithBrackets(age),
		Birthday:       utils.RemoveTextWithBrackets(birthday),
		BloodType:      utils.RemoveTextWithBrackets(bloodtype),
		Bounty:         utils.RemoveTextWithBrackets(bounty),
		Description:    utils.RemoveTextWithBrackets(description),
		DevilFruitName: utils.RemoveTextWithBrackets(devilFruitName),
	}
	resultChan <- character
}

func getCharactersFromLinks(pageLinks []string) []Character {
	var wg sync.WaitGroup
	charactersChan := make(chan Character, len(pageLinks))
	browser := rod.New().MustConnect()
	for _, link := range pageLinks {
		go getCharacterFromPage(browser, link, &wg, charactersChan)
		wg.Add(1)
	}
	go func() {
		wg.Wait()
		close(charactersChan)
	}()

	characters := []Character{}
	for character := range charactersChan {
		characters = append(characters, character)
	}
	return characters

}

func createCharacterCSV(characters []Character) {
	file, err := os.Create("./data/characters.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err.Error())
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	header := []string{"japanese_name", "english_name", "debut", "affiliations", "origin", "age", "birthday", "blood_type", "bounty", "description", "devil_fruit_name"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing row to file:", err.Error())
	}
	for _, character := range characters {
		row := []string{character.JapaneseName, character.EnglishName, character.Debut, character.Affiliations, character.Origin, character.Age, character.Birthday, character.BloodType, character.Bounty, character.Description, character.DevilFruitName}
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing row to file:", err.Error())
		}
	}
	writer.Flush()
}

func LoadCharacters() {
	fmt.Println("Scraping characters...")
	defer utils.Timer("Created characters.csv")()
	characterLinks := getCharacterLinks()
	characters := getCharactersFromLinks(characterLinks)
	createCharacterCSV(characters)
}
