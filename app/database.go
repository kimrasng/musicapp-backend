package app

import (
	"context"
	"database/sql"
	"fmt"
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

	var lastErr error
	for {
		if err := db.PingContext(ctx); err == nil {
			return nil
		} else {
			lastErr = err
		}

		if err := ctx.Err(); err != nil {
			return fmt.Errorf("database did not become ready within 60 seconds: %w (last ping error: %v)", err, lastErr)
		}

		timer := time.NewTimer(time.Second)
		select {
		case <-ctx.Done():
			timer.Stop()
			return fmt.Errorf("database did not become ready within 60 seconds: %w (last ping error: %v)", ctx.Err(), lastErr)
		case <-timer.C:
		}
	}
}
