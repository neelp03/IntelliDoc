package main

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var Db *gorm.DB

// InitDB initializes the database connection using GORM
func InitDB() {
    databaseURL := "db_url_here"
    if databaseURL == "" {
        log.Fatal("DATABASE_URL not set")
    }

    var err error
    Db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Successfully connected to the Timescale Cloud database")
}

// CloseDB closes the database connection.
func CloseDB() {
    sqlDB, err := Db.DB()
    if err != nil {
        log.Fatalf("Failed to get DB from GORM: %v", err)
    }
    sqlDB.Close()
}
