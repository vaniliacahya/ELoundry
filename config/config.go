package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}
	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("DB_NAME")
	defaultConfig.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.Address = os.Getenv("DB_HOST")
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.Port = cnv

	return &defaultConfig
}
