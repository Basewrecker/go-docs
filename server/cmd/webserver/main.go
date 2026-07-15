package main

import (
	"fmt"
	poker "g-test/server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("lets play")
	fmt.Println("type the {name} to record the win")
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
