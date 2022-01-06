package config

import (
	"clean-code/delivery"
	"clean-code/manager"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	Routes         *delivery.Routes
	ApiBaseUrl     string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	// dbEngine := os.Getenv("DB_ENGINE")
	// fmt.Println(dbHost, dbPort, dbName, dbUser, dbPassword)

	infraManager := manager.NewInfra(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManger(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	router := delivery.NewServer(useCaseManager)
	config.Routes = router

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	return config
}
