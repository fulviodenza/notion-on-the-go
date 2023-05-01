package app

import (
	"fmt"
	"log"
	"notiononthego/ent"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	DB_URI = "DB_URI"
)

type App struct {
	*ent.Client
}

func Start() *App {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[godotenv.Load]: failed loading .env file: %v", err.Error())
		return nil
	}

	var dbUri string
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if pair[0] == DB_URI {
			dbUri = pair[1]
		}
	}

	// connect to database with the env db uri
	client, err := ent.SetupAndConnectDatabase(dbUri)
	fmt.Println(client)
	if err != nil {
		log.Fatalf("[ent.SetupAndConnectDatabase]: error in db setup or connection: %v", err.Error())
	}

	return &App{client}
}
