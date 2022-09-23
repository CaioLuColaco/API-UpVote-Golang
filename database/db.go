package database

import (
	"log"

	"github.com/CaioLuColaco/api-upVote-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro na conexão com o Banco de dados")
	}
	DB.AutoMigrate(&models.Coin{})
}
