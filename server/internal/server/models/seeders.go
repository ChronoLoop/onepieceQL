package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gocarina/gocsv"
	"github.com/ikevinws/onepieceQL/internal/server/awsclient"
	"github.com/ikevinws/onepieceQL/pkg/csvmodels"
	"github.com/ikevinws/onepieceQL/pkg/utils"
)

func seedCharacters() {
	if bool, err := CharacterTableExists(); !bool {
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		log.Fatal("Characters table does not exist")
	}

	exePath := utils.GetExePath()
	csvPath := filepath.Join(exePath, "./data/characters.csv")

	charactersCSV, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer charactersCSV.Close()

	characters := []*csvmodels.CharacterCSV{}
	if err := gocsv.UnmarshalFile(charactersCSV, &characters); err != nil {
		panic(err)
	}

	fmt.Println("Seeding characters...")
	var wg sync.WaitGroup

	for _, character := range characters {
		if character.AvatarSrc == "" || !containsWikia(character.AvatarSrc) {
			continue
		}

		go saveCharacter(&wg, character)
		wg.Add(1)
	}
	wg.Wait()

	for _, character := range characters {
		if character.AvatarSrc == "" || containsWikia(character.AvatarSrc) {
			continue
		}
		exists, err := CheckCharacterExists(character)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if !exists {
			if err := CreateCharacter(character); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		}
	}

	fmt.Println("Finished seeding characters")
	saveCSVFile(&characters, csvPath)
}

func seedDevilFruits() {
	if bool, err := DevilFruitTableExists(); !bool {
		if err != nil {
			panic(err)
		}
		log.Fatal("Devil fruits table does not exist")
	}

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

	fmt.Println("Seeding devil fruits...")

	var wg sync.WaitGroup
	for _, devilFruit := range devilFruits {
		if devilFruit.AvatarSrc == "" || !containsWikia(devilFruit.AvatarSrc) {
			continue
		}

		go saveDevilFruit(&wg, devilFruit)
		wg.Add(1)
	}
	wg.Wait()

	for _, devilFruit := range devilFruits {
		if devilFruit.AvatarSrc == "" || containsWikia(devilFruit.AvatarSrc) {
			continue
		}

		exists, err := CheckDevilFruitExists(devilFruit)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if !exists {
			if err := CreateDevilFruit(devilFruit); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		}
	}

	saveCSVFile(&devilFruits, csvPath)
	fmt.Println("Finished seeding devil fruits")
}

func saveCSVFile(in interface{}, csvPath string) {
	tempFilePath := "./data/temp.csv"
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer tempFile.Close()
	if err := gocsv.MarshalFile(in, tempFile); err != nil {
		panic(err)
	}
	err = os.Rename(tempFilePath, csvPath)
	if err != nil {
		fmt.Println("Error replacing the original file:", err)
		return
	}
}

func saveDevilFruit(wg *sync.WaitGroup, devilFruit *csvmodels.DevilFruitCSV) {
	defer wg.Done()
	filename, err := uploadAvatarSrcToS3(devilFruit.AvatarSrc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	devilFruit.AvatarSrc = createS3AvatarSrcString(filename)
}

func saveCharacter(wg *sync.WaitGroup, character *csvmodels.CharacterCSV) {
	defer wg.Done()

	filename, err := uploadAvatarSrcToS3(character.AvatarSrc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	character.AvatarSrc = createS3AvatarSrcString(filename)
}

func uploadAvatarSrcToS3(avatarSrc string) (string, error) {
	AWS_S3_BUCKET := os.Getenv("AWS_S3_BUCKET_NAME")

	resp, err := http.Get(avatarSrc)
	if err != nil {
		errStr := fmt.Sprintf("failed to fetch: %s\n error: %s", avatarSrc, err.Error())
		return "", errors.New(errStr)
	}
	defer resp.Body.Close()

	filename, err := extractImageName(avatarSrc)
	if err != nil {
		errStr := fmt.Sprintf("could not extract file name from avatarSrc: %s", err.Error())
		return "", errors.New(errStr)
	}

	filename = filename + ".jpg"

	_, err = awsclient.S3Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(filename),
	})
	if err != nil {
		fmt.Println("already exists in s3 bucket: ", filename)
		return filename, nil
	}
	uploader := manager.NewUploader(awsclient.S3Client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(AWS_S3_BUCKET),
		Key:         aws.String(filename),
		Body:        resp.Body,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}
	return filename, nil
}

func createS3AvatarSrcString(filename string) string {
	AWS_S3_BUCKET := os.Getenv("AWS_S3_BUCKET_NAME")
	return "https://" + AWS_S3_BUCKET + ".s3.amazonaws.com/" + filename
}

func containsWikia(str string) bool {
	substring := "wikia"
	return strings.Contains(str, substring)
}

func extractImageName(str string) (string, error) {
	pattern := `/([^/]+)\.png`

	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		imageName := matches[1]
		imageName = strings.ReplaceAll(imageName, "_Infobox", "")
		return imageName, nil

	}
	return "", errors.New("could not find image name with this string: " + str)
}

func SeedDatabase() {
	seedDevilFruits()
	seedCharacters()
}
