package main

import (
	"fmt"
	"github.com/qaz027/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("quimbano")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", cfg)
}
