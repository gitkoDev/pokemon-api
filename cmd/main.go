package main

import (
	"os"

	"github.com/gitkoDev/pokemon-db/pkg/repository"
	"github.com/gitkoDev/pokemon-db/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Configuration phase
	if err := initConfigs(); err != nil {
		logrus.Fatal("error loading config files:", err)
	}

	if err := loadEnv(); err != nil {
		logrus.Fatal("error loading .env file:", err)
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
		logrus.Fatalln(err)
	} else {
		logrus.Println("database running on port", dbConfig.Port)
	}

	// Routing phase
	srv := new(server.Server)
	port := viper.GetString("port")
	logrus.Println("server running on port", port)
	if err = srv.Run(port, server.Router(db)); err != nil {
		logrus.Fatalln("error running server", err)
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
