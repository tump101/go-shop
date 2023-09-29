package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func LoadConfig(path string) IConfig{
	envMap,err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v",err)
	}

	return &config{
		app: &app{},
		db: &db{},
		jwt: &jwt{},
	}
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}




type IAppConfig interface{

}
type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

func (c *config) App()IAppConfig{
	return nil
}



type IDbConfig interface{
}
type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}
func (c *config) Db()IDbConfig{
	return nil
}


type IJwtConfig interface{
	
}
type jwt struct {
	secertKey        string
	adminKey         string
	apiKey           string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}
func (c *config) Jwt()IJwtConfig{
	return nil
}