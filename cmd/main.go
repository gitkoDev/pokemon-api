package main

import (
	"os"

	"github.com/gitkoDev/pokemon-db/pkg/handler"
	"github.com/gitkoDev/pokemon-db/pkg/repository"
	"github.com/gitkoDev/pokemon-db/pkg/service"
	"github.com/gitkoDev/pokemon-db/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Configuration phase
	log.SetFormatter(&log.TextFormatter{})

	if err := initConfigs(); err != nil {
		log.Fatal("error loading config files:", err)
	}

	if err := loadEnv(); err != nil {
		log.Fatal("error loading .env file:", err)
	}

	// DB connection phase
	dbConfig := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := repository.NewDB(dbConfig)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("database running on port", dbConfig.Port)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	// Routing phase
	srv := new(server.Server)
	port := viper.GetString("port")
	log.Println("server running on port", port)

	if err = srv.Run(port, handler.InitRoutes()); err != nil {
		log.Fatalln("error running server", err)
	}

}

func initConfigs() error {
	viper.AddConfigPath("configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	return nil
}
