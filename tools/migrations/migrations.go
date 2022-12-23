package main

import (
	"log"
	"migrations/utils"
)

func main() {
	dotenv.Config()
	cmd, path, err := utils.ParseArgs()
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "up":
		err = utils.Run_up(path, db)
	case "down":
		err = utils.Run_down(path, db)
	}

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

