package server

import (
	"github.com/ikevinws/onepieceQL/internal/server/awsclient"
	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/internal/server/models"
	"github.com/joho/godotenv"
)

func StartServer() {
	godotenv.Load("../.env")
	db.Initialize()
	awsclient.ConfigS3()
	models.SeedDatabase()
}
