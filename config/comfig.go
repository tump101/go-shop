package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)



func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRTIE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load write timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load bodylimit failed: %v", err)
				}
				return b
			}(),
			fileLimit: func() int {
				f, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("load filelimit failed: %v", err)
				}
				return f
			}(),
			gcpbucket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnections: func() int {
				m, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load maxconnection failed: %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			secertKey: envMap["JWT_SECRET_KEY"],
			adminKey:  envMap["JWT_ADMIN_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load access expires at failed: %v", err)
				}
				return t
			}(),
			refreshExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refresh expires at failed: %v", err)
				}
				return t
			}(),
		},
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

type IAppConfig interface {
	Url() string //hort:port
	Name() string
	Version() string
	ReadtTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	Gcpbucket() string
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

func (c *config) App() IAppConfig {
	return c.app
}
func (a *app) Url() string                 { return fmt.Sprintf("%s:%d",a.host,a.port) } //hort:port
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadtTimeout() time.Duration { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) Gcpbucket() string           { return a.gcpbucket }

type IDbConfig interface {
	Url() string
	Maxconnections() int
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

func (c *config) Db() IDbConfig {
	return c.db
}
func (b *db) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		b.host, b.port, b.username, b.password, b.database, b.sslMode)
}
func (b *db) Maxconnections() int { return b.maxConnections }

type IJwtConfig interface {
	SecertKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}
type jwt struct {
	secertKey        string
	adminKey         string
	apiKey           string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

func (j *jwt) SecertKey() []byte          { return []byte(j.secertKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
