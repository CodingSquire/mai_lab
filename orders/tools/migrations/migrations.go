package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/innerave/mai_lab/orders/internal/db"
	"github.com/innerave/mai_lab/orders/internal/dotenv"
	"github.com/innerave/mai_lab/orders/tools/migrations/utils"
)

const MAX_RETRIES = 5

func main() {
	dotenv.Config()
	cmd, path, flags, err := utils.ParseArgs()
	if err != nil {
		return
	}

	var dbConn *sql.DB

	if flags.Persist {
		c := make(chan *sql.DB)
		go tryEnshureDB(c)

		dbConn = <-c
		if dbConn == nil {
			log.Fatal("Could not connect to DB due to max retries")
		}
		log.Println("Connected to DB")
	} else {
		dbConn, err = db.Init()
		if err != nil {
			log.Fatal(err)
		}
	}

	switch cmd {
	case "up":
		err = utils.Run_up(path, dbConn)
	case "down":
		err = utils.Run_down(path, dbConn)
	}

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func tryEnshureDB(c chan *sql.DB) {
	for i := 0; i < MAX_RETRIES; i++ {
		dbConn, err := db.Init()
		if err == nil {
			c <- dbConn
			return
		}

		time.Sleep(2 * time.Second)
	}
	c <- nil
}
