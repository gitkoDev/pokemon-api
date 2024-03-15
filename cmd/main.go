package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gitkoDev/pokemon-api/pkg/handler"
	"github.com/gitkoDev/pokemon-api/pkg/repository"
	"github.com/gitkoDev/pokemon-api/pkg/service"
	"github.com/gitkoDev/pokemon-api/server"
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
	hnd := handler.NewHandler(services)

	// Routing phase
	srv := new(server.Server)
	port := viper.GetString("port")

	go func() {
		if err = srv.Run(port, hnd.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalln("error running server", err)
		}
	}()

	log.Println("server running on port", port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()

	log.Println("server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorln("error shutting down server:", err)
	}

	log.Println("database shutting down")
	if err := db.Close(); err != nil {
		log.Errorln("error shutting down database:", err)
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
