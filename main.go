package main

import (
	"fmt"
	"log"

	"github.com/hyperneutr0n/rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read from config: %+v\n", cfg)

	if err = cfg.SetUser("randy"); err != nil {
		log.Fatalf("Couldn't set the current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read from config again: %+v\n", cfg)
}