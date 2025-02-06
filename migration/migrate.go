package main

import (
	"car-rental/database"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
    database.ConnectDB()

    driver, err := postgres.WithInstance(database.DB, &postgres.Config{})
    if err != nil {
        log.Fatal(err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations", // Lokasi folder migrasi
        "postgres",             // Tipe database
        driver,
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Rolling back migrations...")
    err = m.Drop()
    if err != nil {
        log.Fatal(err)
    }
    driver1, err := postgres.WithInstance(database.DB, &postgres.Config{})
    if err != nil {
        log.Fatal(err)
    }

    u, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations", // Lokasi folder migrasi
        "postgres",             // Tipe database
        driver1,
    )
    if err != nil {
        log.Fatal(err)
    }

    err = u.Up()
    if err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }

    fmt.Println("Migration completed successfully!")
}
