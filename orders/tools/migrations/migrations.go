package main

import (
	"database/sql"
	"fmt"
	"log"
	"orders/internal/db"
	"orders/internal/dotenv"
	"os"
	"path/filepath"
)

func main() {
	dotenv.Config()
	cmd, path := parseArgs()

	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "up":
		err = run_up(path, db)
	case "down":
		err = run_down(path, db)
	}

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func parseArgs() (cmd string, path string) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: migration [up|down] <path>")
		os.Exit(1)
	}

	cmd = os.Args[1]
	path = "internal/db/migrations"

	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist", path)
	}

	return
}

func run_down(path string, db *sql.DB) error {
	return run(path, db, "*_down.sql")
}

func run_up(path string, db *sql.DB) error {
	return run(path, db, "*_up.sql")
}

func run(path string, db *sql.DB, glob string) error {
	files, err := filepath.Glob(filepath.Join(path, glob))

	if err != nil {
		return err
	}

	// TODO: sort files

	for _, file := range files {
		fmt.Printf("Running %s\n", file)

		b, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("Error reading file %s: %s", file, err)
		}

		migration := string(b)
		_, err = db.Exec(migration)
		if err != nil {
			return fmt.Errorf("Error running migration %s: %s", file, err)
		}
	}

	return nil
}
