package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luzcn/watchlist-go/src/db"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

var (
	app = kingpin.
		New("go-db-rake", "CLI for database management").
		Version("0.0.1").
		Author("luzcn6")

	// The db connections
	env db.Env

	// db:migrate command
	migrateCmd = app.Command("migrate", "Run database migrations")
	dropCmd    = app.Command("drop", "Delete database")
)

func migrate() {
	fmt.Println("[*] Running DB migration...")

	// migrate the schemas
	env.DB.AutoMigrate(&db.Notes{})
	fmt.Println("[+] Migration complete")
}

func drop() {
	fmt.Println("[*] Deleting DB ...")
	env.DB.DropTable(&db.Notes{})
	fmt.Println("[+] Deleted ...")
}

func connectDB() (err error) {
	dbName := "watchlist-dev"
	conStr := os.Getenv("DATABASE_URL") + "/" + dbName + "?sslmode=disable"

	if os.Getenv("APP_ENV") == "production" {
		dbName = "watchlist"
		conStr = os.Getenv("DATABASE_URL") + "/" + dbName
	}

	env.DB, err = gorm.Open("postgres", conStr)

	if err != nil {
		return err
	}

	env.DB.LogMode(true)
	log.Printf("Connected to database %s\n", conStr)
	return nil
}

func main() {

	// start a db connection
	if err := connectDB(); err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err))
	}
	defer env.DB.Close()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case migrateCmd.FullCommand():
		migrate()
	case dropCmd.FullCommand():
		drop()
	default:
		fmt.Println(`Unknown command, please use "db-rake --help"`)
	}
}
