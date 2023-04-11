package pgsql_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT, status varchar(10) DEFAULT 'active' NOT NULL CHECK (status IN ('active', 'inactive')))")
	db.Exec("CREATE TYPE status AS ENUM ('active', 'inactive')")
	db.Exec("CREATE TABLE IF NOT EXISTS userz (id SERIAL PRIMARY KEY, name TEXT, email TEXT, status status DEFAULT 'active' NOT NULL)")

}
func BenchmarkCreate(t *testing.B) {
	t.Run("Create with constrait check", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			_, err := db.Exec("INSERT INTO users (name, email, status) VALUES ($1, $2, $3)", "John", "john@mail.com", "active")
			if err != nil {
				t.Error(err)
			}
		}
	})
	t.Run("Create with enum", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			_, err := db.Exec("INSERT INTO userz (name, email, status) VALUES ($1, $2, $3)", "John", "john@mail.com", "active")
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Select with constrait check", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			_, err := db.Exec("SELECT * FROM users")
			if err != nil {
				t.Error(err)
			}
		}
	})
	t.Run("Select with enum", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			_, err := db.Exec("SELECT * FROM userz")
			if err != nil {
				t.Error(err)
			}
		}
	})
}
