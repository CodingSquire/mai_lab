package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

type Flags struct {
	Persist bool
}

func ParseArgs() (string, string, Flags, error) {
	flags := Flags{}
	if len(os.Args) < 2 {
		fmt.Println("Usage: migration [up|down] <path> [--persist]")
		return "", "", flags, fmt.Errorf("Not enough arguments")
	}

	cmd := os.Args[1]
	path := "internal/db/migrations"

	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	if len(os.Args) > 3 && os.Args[3] == "--persist" {
		flags.Persist = true
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", "", flags, fmt.Errorf("Path %s does not exist", path)
	}

	return cmd, path, flags, nil
}

func Run_down(path string, db *sql.DB) error {
	return Run(path, db, "*_down.sql")
}

func Run_up(path string, db *sql.DB) error {
	return Run(path, db, "*_up.sql")
}

func Run(path string, db *sql.DB, glob string) error {
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
