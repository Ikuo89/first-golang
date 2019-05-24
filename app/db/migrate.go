package main

import (
  "fmt"
  "os"
  "database/sql"
  "github.com/rubenv/sql-migrate"
  _ "github.com/lib/pq"
)

func main() {
  host := os.Getenv("DB_HOST")
  user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  dbname := os.Getenv("DB_NAME")
  db, err := sql.Open("postgres", "host="+host+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  migrations := &migrate.FileMigrationSource{
    Dir: "db/migrations",
  }

  n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
  if err != nil {
    panic(err.Error())
  }

  fmt.Printf("Applied %d migrations!\n", n)
}
