package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    InitDB()
    defer CloseDB()

    // Seed the database with sample data
    seedData()

    router := gin.Default()
    router.POST("/search", SearchHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Starting server on port %s...", port)
    router.Run(":" + port)
}
