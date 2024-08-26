package shared

import (
    "fmt"
    // "log"
    // "os"

    // "github.com/joho/godotenv"
    // "gorm.io/driver/postgres"
    // "gorm.io/gorm"
)

// var DB *gorm.DB

func DbInit() {
    // err := godotenv.Load() // Load .env file if available
    // if err != nil {
    //     log.Printf("Error loading .env file: %v", err)
    // }

    // dbHost := os.Getenv("DB_HOST")
    // dbPort := os.Getenv("DB_PORT")
    // dbUser := os.Getenv("DB_USER")
    // dbPass := os.Getenv("DB_PASS")
    // dbName := os.Getenv("DB_NAME")

    // dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    //     dbHost, dbUser, dbPass, dbName, dbPort)

    // DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    // if err != nil {
    //     log.Fatalf("Failed to connect to database: %v", err)
    // }

    fmt.Println("Database connection successful")
}

// func GetDB() *gorm.DB {
//     return DB
// }
