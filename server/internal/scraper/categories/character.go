package categories

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/gocarina/gocsv"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

var CHARACTER_LIST_LINKS = [...]string{
	"https://onepiece.fandom.com/wiki/List_of_Canon_Characters?action=render",
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
		wg.Add(1)
		go getCharacterLinksFromPage(link, &wg, linksChan)
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

func getCharacterFromPage(browser *rod.Browser, pageLink string, wg *sync.WaitGroup, resultChan chan<- csvmodels.CharacterCSV) {

	defer wg.Done()
	page := browser.MustPage(pageLink)
	var asideElement *rod.Element = nil
	rod.Try(func() {
		asideElement = page.MustElement("body > div > aside")
	})
	if asideElement == nil {
		fmt.Printf("could not find aside element: %s\n", pageLink)
		resultChan <- csvmodels.CharacterCSV{}
		return
	}
	var characterSection *rod.Element = nil
	sectionIndex := 1
	rod.Try(func() {
		characterSection = asideElement.MustElement("section:nth-of-type(" + strconv.Itoa(sectionIndex) + ")")
	})
	if characterSection == nil {
		fmt.Printf("could not find character section element: %s\n", pageLink)
		resultChan <- csvmodels.CharacterCSV{}
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
	avatarSrc := ""

	rod.Try(func() {
		japaneseName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/japanese name:/i").MustNext().MustText()
	})
	rod.Try(func() {
		englishName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/english name:/i").MustNext().MustText()
	})
	if englishName == "N/A" {
		rod.Try(func() {
			englishName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/romanized name:/i").MustNext().MustText()
		})
	}

	//check next section if name does not exist
	if englishName == "N/A" && japaneseName == "N/A" {
		sectionIndex++
		rod.Try(func() {
			characterSection = asideElement.MustElement("section:nth-of-type(" + strconv.Itoa(sectionIndex) + ")")
		})

		rod.Try(func() {
			japaneseName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/japanese name:/i").MustNext().MustText()
		})
		rod.Try(func() {
			englishName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/english name:/i").MustNext().MustText()
		})
		if englishName == "N/A" {
			rod.Try(func() {
				englishName = characterSection.Timeout(1*time.Second).MustElementR("h3", "/romanized name:/i").MustNext().MustText()
			})
		}

	}

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
	rod.Try(func() {
		avatarImgElement := asideElement.Timeout(1 * time.Second).MustElement("div > div.wds-tab__content.wds-is-current > figure > a > img")
		imgSrc := avatarImgElement.MustAttribute("src")
		avatarSrc = *imgSrc
	})
	if avatarSrc == "" {
		rod.Try(func() {
			avatarImgElement := asideElement.Timeout(1 * time.Second).MustElement("figure > a > img")
			imgSrc := avatarImgElement.MustAttribute("src")
			avatarSrc = *imgSrc
		})
	}

	for i := 1; i < 6; i++ {
		selector := "body > div > p:nth-of-type(" + strconv.Itoa(i) + ")"
		rod.Try(func() {
			if descriptionParagraphElement := page.Timeout(1 * time.Second).MustElement(selector); descriptionParagraphElement != nil {
				description = descriptionParagraphElement.MustText()
			}
		})
		if description != "N/A" && description != "" {
			break
		}
	}

	if description == "" || description == "N/A" {
		rod.Try(func() {

			descriptionStr := page.Timeout(2 * time.Second).MustEval(`()=>{
                const aside = document.querySelector('body > div > aside');

                const endParagraph = document.querySelector('body > div> p:nth-of-type(2)');

                let description = "";
                let currentElement = aside.nextSibling;

                while (currentElement && currentElement !== endParagraph) {
                    description += currentElement.textContent
                    currentElement = currentElement.nextSibling;
                }
                return description
            }`).Str()
			if descriptionStr != "" {
				description = descriptionStr
			}
		})
	}

	if description == "" {
		description = "N/A"
	}

	devilFruitName := "N/A"
	rod.Try(func() {
		characterDevilFruitSection := asideElement.Timeout(1 * time.Second).MustElement("section:nth-of-type(" + strconv.Itoa(sectionIndex+1) + ")")
		devilFruitName = characterDevilFruitSection.Timeout(1*time.Second).MustElementR("h3", "/japanese name:/i").MustNext().MustText()
	})

	character := csvmodels.CharacterCSV{
		JapaneseName:   processCategoryString(japaneseName),
		EnglishName:    processCategoryString(englishName),
		Debut:          processCategoryString(debut),
		Affiliations:   processCategoryString(affiliations),
		Origin:         processCategoryString(origin),
		Age:            processCategoryString(age),
		Birthday:       processCategoryString(birthday),
		BloodType:      processCategoryString(bloodtype),
		Bounty:         processCategoryString(bounty),
		Description:    processCategoryString(description),
		DevilFruitName: processCategoryString(devilFruitName),
		AvatarSrc:      avatarSrc,
		URL:            pageLink,
	}
	resultChan <- character
}

func getCharactersFromLinks(pageLinks []string) []csvmodels.CharacterCSV {
	var wg sync.WaitGroup
	charactersChan := make(chan csvmodels.CharacterCSV, len(pageLinks))
	browser := rod.New().MustConnect()
	for _, link := range pageLinks {
		wg.Add(1)
		go getCharacterFromPage(browser, link, &wg, charactersChan)
	}
	go func() {
		wg.Wait()
		close(charactersChan)
	}()

	characters := []csvmodels.CharacterCSV{}
	for character := range charactersChan {
		characters = append(characters, character)
	}
	return characters

}

func createCharacterCSV(characters []csvmodels.CharacterCSV) {
	file, err := os.Create("./data/characters.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	if err := gocsv.MarshalFile(&characters, file); err != nil {
		fmt.Println("Error writing rows to file:", err)
	}
}

func ScrapeCharacters() {
	fmt.Println("Scraping characters...")
	defer utils.Timer("Created characters.csv")()
	characterLinks := getCharacterLinks()
	characters := getCharactersFromLinks(characterLinks)
	createCharacterCSV(characters)
	CheckCharacters()
}
