package config

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConfig() *gorm.DB {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to load .env", err)
	}

	host := viper.GetString("HOST")
	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbname := viper.GetString("DBNAME")
	port := viper.GetString("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("invalid to connect database", err)
		os.Exit(2)
	}

	log.Println("connect db successfully")
	return db
}