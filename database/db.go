package database

import (
	"log"
	"os"

	"github.com/CaioLuColaco/api-upVote-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Panic("Erro na conexão com o Banco de dados")
	}
	DB.AutoMigrate(&models.Coin{})
}
