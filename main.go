package main

import (
	"log"

	"musicapp-backend/app"
)

func main() {
	if err := app.LoadEnv(); err != nil {
		log.Printf(".env not loaded: %v", err)
	}

	db, err := app.OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	address := ":" + app.GetEnv("PORT", "8080")
	log.Printf("music API listening on %s", address)
	log.Fatal(app.ListenAndServe(db, address))
}
