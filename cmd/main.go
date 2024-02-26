package main

import (
	"log"

	"github.com/gitkoDev/pokemon-db/pkg/repository"
	"github.com/gitkoDev/pokemon-db/server"
	"github.com/spf13/viper"
)

func main() {
	// Configuration phase
	if err := initConfigs(); err != nil {
		log.Fatalln("error loading config files:", err)
	}

	// DB connection phase

	dbConfig := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := repository.NewDB(dbConfig)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("database running on port", dbConfig.Port)
	}

	// Routing phase
	srv := new(server.Server)
	port := viper.GetString("port")
	log.Println("server running on port", port)
	if err = srv.Run(port, server.Router(db)); err != nil {
		log.Fatalln("error running server", err)
	}

}

func initConfigs() error {
	viper.AddConfigPath("configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
