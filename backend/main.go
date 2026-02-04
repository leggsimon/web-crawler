package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/internal/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	dsn := getDSN()
	port := getPort()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping db: %v", err)
	}

	sourceDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		log.Fatalf("migrate source: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", sourceDriver, "mysql://"+dsn)
	if err != nil {
		log.Fatalf("migrate: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrate up: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	mux.Handle("/ready", handlers.Ready(db))

	mux.HandleFunc("/urls", handlers.URLs(db))

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}

func getDSN() string {
	if dsn := os.Getenv("MYSQL_DSN"); dsn != "" {
		return dsn
	}
	user := getEnv("MYSQL_USER", "root")
	password := os.Getenv("MYSQL_PASSWORD")
	host := getEnv("MYSQL_HOST", "localhost")
	port := getEnv("MYSQL_PORT", "3306")
	database := getEnv("MYSQL_DATABASE", "app")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}

func getPort() string {
	return getEnv("PORT", "8080")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
