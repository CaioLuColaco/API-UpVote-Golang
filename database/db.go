package database

import (
	"log"
	"os"

	"github.com/CaioLuColaco/api-upVote-golang/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func ConnectToDatabase() {
	DB, err = gorm.Open(postgres.Open(goDotEnvVariable("DATABASE_URL")))
	if err != nil {
		log.Panic("Erro na conexão com o Banco de dados")
	}
	DB.AutoMigrate(&models.Coin{})
}
