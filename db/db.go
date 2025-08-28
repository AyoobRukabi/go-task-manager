package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "github.com/AyoobRukabi/go-task-manager/models"
)


var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=Go6969 dbname=task_manager port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    database.AutoMigrate(&models.Task{})

    DB = database
    fmt.Println("Database connected successfully!")
}
