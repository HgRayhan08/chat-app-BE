package connection

import (
	"chat-app/internal/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection(config config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)

	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", dsn)

		if err != nil {
			log.Printf("Attempt %d: failed to open DB connection %v\n", i, err)
		} else {
			err = db.Ping()
			if err == nil {
				log.Println("✅ Database Connected Successfully")
				return db
			}
			log.Printf("Attempt %d: Database not ready yet: %v\n", i, err)
		}
		log.Println("Waiting for Database!!!...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Cannot connection to database: ", err)
	return nil
}
