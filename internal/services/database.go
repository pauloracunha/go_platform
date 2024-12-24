package services

import (
	"fmt"
	"log"
	"portal/config"
	"portal/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "3306")
	user := config.GetEnv("DB_USER", "root")
	password := config.GetEnv("DB_PASSWORD", "root")
	dbname := config.GetEnv("DB_NAME", "portal_ong")

	// Configuração da string de conexão do MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados MySQL: %v", err)
	}

	// Migração automática
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Erro ao migrar modelo User: %v", err)
	}
	log.Println("Banco de dados conectado e migrado com sucesso!")
}
