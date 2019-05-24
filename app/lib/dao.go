package lib

import (
  "fmt"
  "os"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"

  "github.com/projects/app/model"
)

var db = NewDbConnection()

func NewDbConnection() *gorm.DB {
  host := os.Getenv("DB_HOST")
  user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  dbname := os.Getenv("DB_NAME")
  db, err := gorm.Open("postgres", "host="+host+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
  if err != nil {
    fmt.Println(err.Error())
    panic(err.Error())
  }

  db.AutoMigrate(model.User{})
  db.LogMode(true)

  return db
}

func DbConnection() *gorm.DB {
  return db
}
