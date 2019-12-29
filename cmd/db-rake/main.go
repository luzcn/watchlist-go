package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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
	env.DB.AutoMigrate(&db.Note{}, &db.Product{})

	// add the foreign key
	env.DB.Model(&db.Note{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

	fmt.Println("[+] Migration complete")
}

func drop() {
	fmt.Println("[*] Deleting DB ...")
	env.DB.DropTable(&db.Note{})
	fmt.Println("[+] Deleted ...")
}

func connectDB() error {

	var conStr string
	if os.Getenv("APP_ENV") == "production" {
		conStr = os.Getenv("DATABASE_URL") + "/watchlist"
	} else {
		conStr = os.Getenv("DATABASE_URL") + "/watchlist-dev?sslmode=disable"
	}

	dbCon, err := gorm.Open("postgres", conStr)

	if err != nil {
		return err
	}

	dbCon.LogMode(true)
	env = db.Env{DB: dbCon}

	log.Printf("Connected to database %s\n", conStr)
	return nil
}

func main() {
	// load environment variables
	_ = godotenv.Load()

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
