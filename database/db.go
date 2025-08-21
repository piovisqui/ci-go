package database

import (
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	endereco := os.Getenv("HOST")
	usuario := os.Getenv("USER")
	senha := os.Getenv("PASSWORD")
	nomeBanco := os.Getenv("DBNAME")
	portaBanco := os.Getenv("PORT")
	//stringDeConexao := "host=" + endereco + " user=" + usuario + " password=" + senha + " dbname=" + nomeBanco + " port=5432 sslmode=disable"
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados - ")
	} else {
		log.Println("Conexão com banco de dados realizada com sucesso")
	}

	DB.AutoMigrate(&models.Aluno{})
}
