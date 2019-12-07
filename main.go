package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luzcn/watchlist-go/src/db"
	"github.com/luzcn/watchlist-go/src/web/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

var env db.Env

func main() {

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "5000"
	}

	// connect to db
	dbName := "watchlist-dev"
	conStr := os.Getenv("DATABASE_URL") + "/" + dbName + "?sslmode=disable"

	if os.Getenv("APP_ENV") == "production" {
		dbName = "watchlist"
		conStr = os.Getenv("DATABASE_URL") + "/" + dbName
	}

	// start a db connection
	var err error
	env.DB, err = gorm.Open("postgres", conStr)
	defer env.DB.Close()

	if err != nil {
		panic(err)
	}

	// register web handler
	r := mux.NewRouter()
	handlers.Load(r, &env)

	server := &http.Server{
		Addr:         ":" + PORT,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      r,
	}
	log.Printf("action=start-server msg=\"Listening on port %s\"", PORT)
	_ = server.ListenAndServe()
}
