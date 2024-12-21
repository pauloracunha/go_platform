package services

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portal.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	log.Println("Banco de dados conectado com sucesso")
}
