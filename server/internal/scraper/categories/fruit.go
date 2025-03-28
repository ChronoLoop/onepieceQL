package categories

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/gocarina/gocsv"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

var DEVIL_FRUIT_CATEGORY_LINKS = [...]string{
	"https://onepiece.fandom.com/wiki/Category:Paramecia?action=render",
	"https://onepiece.fandom.com/wiki/Category:Logia?action=render",
	"https://onepiece.fandom.com/wiki/Category:Zoan?action=render",
	"https://onepiece.fandom.com/wiki/Category:Mythical_Zoan?action=render",
	"https://onepiece.fandom.com/wiki/Category:Ancient_Zoan?action=render",
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
		wg.Add(1)
		go getDevilFruitLinksFromPage(link, &wg, linksChan)
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

func getDevilFruitFromPage(browser *rod.Browser, pageLink string, wg *sync.WaitGroup, resultChan chan<- csvmodels.DevilFruitCSV) {
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
	previousOwner := "N/A"
	rod.Try(func() {
		if previousOwnerHeader := asideElement.Timeout(2*time.Second).MustElementR("section h3", "/previous (user|owner):/i"); previousOwnerHeader != nil {
			previousOwner = previousOwnerHeader.MustNext().MustText()
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
	for i := 1; i < 6; i++ {
		selector := "body > div > p:nth-of-type(" + strconv.Itoa(i) + ")"
		rod.Try(func() {
			if descriptionParagraphElement := page.Timeout(2 * time.Second).MustElement(selector); descriptionParagraphElement != nil {
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

	avatarSrc := ""
	rod.Try(func() {
		avatarImgElement := asideElement.Timeout(1 * time.Second).MustElement("figure > a > img")
		imgSrc := avatarImgElement.MustAttribute("src")
		avatarSrc = *imgSrc
	})

	devilFruit := csvmodels.DevilFruitCSV{
		JapaneseName: processCategoryString(japaneseName),
		EnglishName:  processFruitEnglishName(englishName),
		Meaning:      processCategoryString(meaning),
		UsageDebut:   processCategoryString(usageDebut),
		Type:         processCategoryString(fruitType),

		PreviousOwner: processCategoryString(previousOwner),
		CurrentOwner:  processCategoryString(currentUser),
		Description:   processCategoryString(description),
		AvatarSrc:     avatarSrc,
		URL:           pageLink,
	}
	resultChan <- devilFruit
}
func getDevilFruitsFromLinks(pageLinks []string) []csvmodels.DevilFruitCSV {
	devilFruits := []csvmodels.DevilFruitCSV{}
	var wg sync.WaitGroup
	devilFruitChan := make(chan csvmodels.DevilFruitCSV, len(pageLinks))
	browser := rod.New().MustConnect()
	for _, link := range pageLinks {
		wg.Add(1)
		go getDevilFruitFromPage(browser, link, &wg, devilFruitChan)
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

func createDevilFruitCsv(devilFruits []csvmodels.DevilFruitCSV) {
	file, err := os.Create("./data/fruits.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	if err := gocsv.MarshalFile(&devilFruits, file); err != nil {
		fmt.Println("Error writing rows to file:", err)
	}
}

func ScrapeFruits() {
	fmt.Println("Scraping devil fruits...")
	defer utils.Timer("Created fruits.csv")()
	devilFruitLinks := getDevilFruitLinks()
	devilFruits := getDevilFruitsFromLinks(devilFruitLinks)
	createDevilFruitCsv(devilFruits)
	CheckFruits()
}
