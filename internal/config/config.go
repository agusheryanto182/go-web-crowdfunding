package config

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort int
	Secret  string

	CCName      string
	CCAPIKey    string
	CCAPISecret string
	CCFolder    string

	Database database
	Redis    Redis

	Midtrans Midtrans

	OpenAiApiKey string
}

type database struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type Redis struct {
	Addr string
	Pass string
}

type Midtrans struct {
	ClientKeyMidtrans string
	ServerKeyMidtrans string
}

func InitConfig() *Config {
	return loadConfig()
}

func loadConfig() *Config {
	var res = new(Config)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	if value, found := os.LookupEnv("PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.AppPort = port
	}

	if value, found := os.LookupEnv("SECRET"); found {
		res.Secret = value
	}

	if value, found := os.LookupEnv("CCNAME"); found {
		res.CCName = value
	}

	if value, found := os.LookupEnv("CCAPIKEY"); found {
		res.CCAPIKey = value
	}
	if value, found := os.LookupEnv("CCAPISECRET"); found {
		res.CCAPISecret = value
	}
	if value, found := os.LookupEnv("CCFOLDER"); found {
		res.CCFolder = value
	}

	if value, found := os.LookupEnv("DBHOST"); found {
		res.Database.DbHost = value
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		res.Database.DbPort = value
	}

	if value, found := os.LookupEnv("DBUSER"); found {
		res.Database.DbUser = value
	}

	if value, found := os.LookupEnv("DBPASS"); found {
		res.Database.DbPass = value
	}

	if value, found := os.LookupEnv("DBNAME"); found {
		res.Database.DbName = value
	}

	if value, found := os.LookupEnv("REDIS_ADDR"); found {
		res.Redis.Addr = value
	}

	if value, found := os.LookupEnv("REDIS_PASS"); found {
		res.Redis.Pass = value
	}

	if value, found := os.LookupEnv("CLIENT_MIDTRANS"); found {
		res.Midtrans.ClientKeyMidtrans = value
	}

	if value, found := os.LookupEnv("SERVER_MIDTRANS"); found {
		res.Midtrans.ServerKeyMidtrans = value
	}

	if value, found := os.LookupEnv("OPENAIAPIKEY"); found {
		res.OpenAiApiKey = value
	}

	return res
}

func BootConfig() *Config {
	return loadConfig()
}
