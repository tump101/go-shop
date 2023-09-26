package config

import (
	"os"
	"time"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type app struct {
	host         string
	port         int
	name         string
	verion       string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

type db struct {
	db             string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	maxConnections int
}

type jwt struct {
	adminKey         string
	secerKey         string
	apiKey           string
	accessExpiresAt  int //bytes
	refreshExpiresAt int //bytes
}
