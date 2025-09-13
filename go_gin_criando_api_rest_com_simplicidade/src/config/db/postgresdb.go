package postgresdb

import (
	"go_gin_criando_api_rest_com_simplicidade/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Student{})
}
