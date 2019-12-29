package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/luzcn/watchlist-go/src/db"
	"github.com/luzcn/watchlist-go/src/web/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	env  db.Env
	deps handlers.Deps
)

// connect to db
func connectDB() error {
	var conStr string
	if os.Getenv("APP_ENV") == "production" {
		conStr = os.Getenv("DATABASE_URL") + "/watchlist"
	} else {
		conStr = os.Getenv("DATABASE_URL") + "/watchlist-dev" + "?sslmode=disable"
	}

	dbCon, err := gorm.Open("postgres", conStr)

	if err != nil {
		return err
	}

	dbCon.LogMode(true)
	env = db.Env{DB: dbCon}

	deps = handlers.Deps{DB: &env}

	log.Printf("Connected to database %s\n", conStr)

	return nil
}
func main() {

	// load environment variable
	_ = godotenv.Load()

	// connect to DB
	if err := connectDB(); err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err))
	}
	defer env.DB.Close()

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "5000"
	}

	// register web handler
	r := mux.NewRouter()
	handlers.Load(r, &deps)

	server := &http.Server{
		Addr:         ":" + PORT,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      r,
	}
	log.Printf("action=start-server msg=\"Listening on port %s\"", PORT)
	_ = server.ListenAndServe()
}
