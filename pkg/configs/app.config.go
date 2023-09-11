package configs

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	StageStatus string

	DbType        string
	DbHost        string
	DbPort        string
	DbPassword    string
	DbUser        string
	DbName        string
	JwtSecret     string
	ResultBackend string

	BrokerUrl string
)

func convertEnvToInt(key string) int {
	valueStr := os.Getenv(key)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatal("ERROR: "+key+": "+valueStr+" ", err)
	}
	return valueInt
}

func init() {
	StageStatus = os.Getenv("STAGE_STATUS")
	JwtSecret = os.Getenv("JWT_SECRET_KEY")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	BrokerUrl = os.Getenv("BROKER_URL")
	ResultBackend = os.Getenv("RESULT_BACKEND")
}

type TransactOptsConfig struct {
	PrivateKey string `json:"private_key"`
	ChainID    int64  `json:"chain_id"`
}
