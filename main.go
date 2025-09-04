package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hyperneutr0n/rss-aggregator/internal/config"
	"github.com/hyperneutr0n/rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		log.Fatalf("Failed connecting to database")
	}
	dbQueries := database.New(db)

	s := &state{dbQueries, &cfg}

	cmds := commands{make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	if err := cmds.run(s, command{cmdName, cmdArgs}); err != nil {
		log.Fatal(err);
	}
}