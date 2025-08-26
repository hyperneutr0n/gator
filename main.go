package main

import (
	"log"
	"os"

	"github.com/hyperneutr0n/rss-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	s := &state{&cfg}

	cmds := commands{make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	if err := cmds.run(s, command{cmdName, cmdArgs}); err != nil {
		log.Fatal(err);
	}
}