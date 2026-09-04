package app

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDatabase() (*sql.DB, error) {
	dsn := GetEnv("MYSQL_DSN", "musicapp:musicapp@tcp(mysql:3306)/musicapp?parseTime=true")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := waitForDatabase(db); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func waitForDatabase(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for {
		if err := db.PingContext(ctx); err == nil {
			return nil
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return errors.New("database did not become ready within 60 seconds")
		}
		time.Sleep(time.Second)
	}
}
