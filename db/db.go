package db

import (
    "fmt"
    "log"
    "github.com/AyoobRukabi/go-task-manager/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=Go6969 dbname=task_manager port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("[error] failed to connect to database: %v", err)
    }

    // Auto migrate tables
    err = DB.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatalf("[error] failed to migrate database: %v", err)
    }

    fmt.Println("Database connected and migrated successfully!")
}
